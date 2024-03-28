package ws

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/websocket"
	"log"
	"strconv"
	"videoweb/biz/utils"
	"videoweb/database/DB/dao"
	"videoweb/pkg/e"
	"videoweb/rabbitmq"
)

// 处理客户端发送的消息
func (c *Client) read() {

	//Register <- c
	defer func() {
		Unregister <- c
		c.conn.Close()
	}()
	for {

		msg := new(Message)
		if c.conn == nil {
			log.Println("c.conn is nil, cannot read from it")
			return
		}
		err := c.conn.ReadJSON(&msg)
		if err != nil {

			log.Println("your input is not suitable")

			log.Println("err = ", err)
			c.conn.Close()
			break
		}

		msgToBroadcast := new(MsgFromBroadcast)
		msgToBroadcast = &MsgFromBroadcast{
			Type:    msg.Type,
			Content: msg.Content,
			Target:  c.Target,
			FromUid: c.FromUid,
		}

		c.Type = msg.Type

		if len(msg.Content) > e.MaxStore {
			c.conn.WriteMessage(websocket.TextMessage, []byte("消息过长"))
		}
		//2 获取与某人的全部历史记录
		//3 获取与某人有关的未读记录
		//4 获取群聊的历史记录
		msgDao := dao.GetMsgDao(context.Background())

		switch msg.Type {

		case e.GetHistoryFromSingleChat:
			//log.Println("success")
			msgs, err := msgDao.GetHistoryFromSingleChat(msg.PageNum, c.FromUid, c.Target)
			if err != nil || len(msgs) == 0 {
				c.conn.WriteMessage(websocket.TextMessage, []byte("查找不到相关记录"))
			} else {
				c.SendHistory(msgs)

			}

		case e.GetUnreadFromSingleChat:
			msgs, err := msgDao.GetHistoryFromGroupChat(msg.PageNum, c.Target)
			if err != nil || len(msgs) == 0 {
				c.conn.WriteMessage(websocket.TextMessage, []byte("查找不到相关记录"))
			} else {
				c.SendHistory(msgs)
			}

		case e.GetHistoryFromGroupChat:
			msgs, err := msgDao.GetUnreadFromSingleChat(msg.PageNum, c.FromUid, c.Target)
			if err != nil || len(msgs) == 0 {
				c.conn.WriteMessage(websocket.TextMessage, []byte("查找不到相关记录"))
			} else {
				c.SendHistory(msgs)
			}

		}
		// 广播消息给客户端

		Broadcast <- msgToBroadcast

	}
}

// 向客户端发送消息
func (c *Client) write() {
	defer func() {

		c.conn.Close()
	}()

	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}

}

var (
	ClientsSingleMap = make(map[string]*Client)
	ClientsMap       = make(map[string]map[string]*Client) //前者群聊id，后者用户id
	Clients          = make(map[*Client]bool)              // 存储所有客户端
	Broadcast        = make(chan *MsgFromBroadcast)        // 广播消息通道
	Register         = make(chan *Client)
	Unregister       = make(chan *Client)
	upgrader         = websocket.HertzUpgrader{}
)

// 处理WebSocket连接
func HandleConnections(c context.Context, ctx *app.RequestContext) { //(w http.ResponseWriter, r *http.Request) {
	claim, err := utils.ParseToken(ctx.Query("token"))
	if err != nil {

		return
	}
	log.Println("uid = ", claim.Uid)
	//升级HTTP连接为WebSocket连接
	err = upgrader.Upgrade(ctx, func(conn *websocket.Conn) {

		client := newClient(conn, strconv.FormatInt(claim.Uid, 10))

		if len(ctx.Query("to_uid")) != 0 {
			client.Target = ctx.Query("to_uid")
			client.Type = e.SingleChat
		} else {
			client.Target = ctx.Query("group_id")
			client.Type = e.GroupChat
		}

		Register <- client

		list, err := rabbitmq.Consume(client.FromUid)
		if err != nil {
			//conn.Close()
			return
		}

		for _, v := range list {
			client.send <- v
		}

		msgDao := dao.GetMsgDao(context.Background())
		err = msgDao.TurnToRead(client.Target, client.FromUid)
		if err != nil {
			//conn.Close()
			return
		}
		// 启动goroutine处理客户端消息
		go client.write()
		client.read()

	})

	if err != nil {
		log.Println(err)
		return
	}
}
