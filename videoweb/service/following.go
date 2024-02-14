package service

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"strconv"
	"sync"
	"videoweb/biz/model/hertz/following"
	"videoweb/database/DB/dao"
	"videoweb/database/DB/model"
	"videoweb/response"
)

func ListFollowing(c context.Context, ctx *app.RequestContext, req *following.FollowingListReq) (interface{}, error) {
	//claim, err := utils.ParseToken(string(ctx.GetHeader("access_token")))
	//if err != nil {
	//	return response.BadResponse(), err
	//}禁止自己关注自己
	uid, _ := strconv.ParseInt(req.GetUid(), 10, 64)
	var res []model.Relation
	var count int64
	if req.GetPageNum() >= 0 && req.GetPageSize() >= 0 {
		err := dao.Db.Model(&model.Relation{}).Offset(int(req.GetPageNum()*req.GetPageSize())).
			Limit(int(req.GetPageSize())).Where("from_uid = ? AND status = ?", uid, 0).
			Find(&res).Count(&count).Error
		if err != nil {
			return response.BadResponse(), err
		}
	} else {
		return response.BadResponse(), errors.New("请重新操作")
	}
	//else if  req.GetPageSize() == 0 {
	//	err := dao.Db.Model(&model.Relation{}).Where("from_uid = ?", uid).
	//		Find(&res).Count(&count).Error
	//	if err != nil {
	//		return response.BadResponse(), err
	//	}}
	resultCh := make(chan *model.User)
	var users []*model.User
	var wg sync.WaitGroup
	for i := 0; i < len(res); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var item *model.User
			err := dao.Db.Model(&model.User{}).Where("uid = ?", res[i].ToUid).Find(&item).Error
			if err == nil {
				resultCh <- item
				//users = append(users, item)
			}
		}(i)
	}
	go func() {
		wg.Wait()       // 等待所有 goroutine 执行完成
		close(resultCh) // 关闭通道
	}()

	// 从通道中读取结果
	for user := range resultCh {
		users = append(users, user)
	}
	//wg.Wait()
	return response.UserInfoResponse(users, count), nil
}
