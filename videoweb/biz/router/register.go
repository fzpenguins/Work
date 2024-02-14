// Code generated by hertz generator. DO NOT EDIT.

package router

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	comment "videoweb/biz/router/comment"
	common "videoweb/biz/router/common"
	follower "videoweb/biz/router/follower"
	following "videoweb/biz/router/following"
	friends "videoweb/biz/router/friends"
	like "videoweb/biz/router/like"
	relation "videoweb/biz/router/relation"
	user "videoweb/biz/router/user"
	video "videoweb/biz/router/video"
)

// GeneratedRegister registers routers generated by IDL.
func GeneratedRegister(r *server.Hertz) {
	//INSERT_POINT: DO NOT DELETE THIS LINE!
	relation.Register(r)

	comment.Register(r)

	follower.Register(r)

	following.Register(r)

	friends.Register(r)

	like.Register(r)

	video.Register(r)

	user.Register(r)

	common.Register(r)
}