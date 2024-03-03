package service

import (
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"
	"strconv"
	"time"
	"videoweb/biz/model/hertz/comment"
	"videoweb/biz/model/hertz/common"
	"videoweb/biz/utils"
	"videoweb/database/DB/dao"
	"videoweb/database/DB/model"
	"videoweb/database/cache"
	"videoweb/pkg/e"
	"videoweb/response"
)

func PublishComment(c context.Context, ctx *app.RequestContext, commentReq *comment.CommentPublishReq) (interface{}, error) {
	vid, _ := strconv.ParseInt(commentReq.GetVid(), 10, 64)

	if len(commentReq.GetContent()) == 0 {
		return response.BadResponse(), errors.New("评论不能为空")
	}
	claim, err := utils.ParseToken(string(ctx.GetHeader("access_token")))
	if err != nil {
		return response.BadResponse(), err
	}
	var v model.Video
	err = dao.Db.Model(&model.Video{}).Where("vid = ?", vid).First(&v).UpdateColumn("comment_count", gorm.Expr("comment_count+?", 1)).Error
	if err != nil {
		return response.BadResponse(), err
	}

	com := model.Comment{
		Uid:        claim.Uid,
		Vid:        vid, //commentReq.GetVid(),
		Cid:        0,
		ParentId:   0,
		LikeCount:  0,
		ChildCount: 0,
		Content:    commentReq.Content,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
	}
	err = dao.Db.Model(&model.Comment{}).Create(&com).Error
	if err != nil {
		return response.BadResponse(), err
	}
	return response.GoodResponse(), nil
}

func ListComment(c context.Context, ctx *app.RequestContext, commentReq *comment.CommentListReq) (interface{}, error) {
	vid, _ := strconv.ParseInt(commentReq.GetVid(), 10, 64)

	var comments []model.Comment
	err := dao.Db.Model(&model.Comment{}).Where("vid = ?", vid).Limit(int(commentReq.GetPageSize())).
		Offset(int(commentReq.GetPageNum() * commentReq.GetPageSize())).
		Find(&comments).Error
	if err != nil {
		return response.BadResponse(), err
	}
	resp := new(comment.CommentResponse)
	var res []*comment.CommentItems
	for _, v := range comments {
		com := new(comment.CommentItems)
		com = &comment.CommentItems{
			Uid:        strconv.FormatInt(v.Uid, 10),
			Vid:        strconv.FormatInt(v.Vid, 10),
			Cid:        strconv.FormatInt(v.Cid, 10),
			ParentId:   strconv.FormatInt(v.ParentId, 10),
			LikeCount:  cache.LikeCount(c, v.Vid),
			ChildCount: 0, //暂时不对评论作评论
			Content:    v.Content,
			CreatedAt:  v.CreatedAt,
			UpdatedAt:  v.UpdatedAt,
			DeletedAt:  v.DeletedAt,
		}

		res = append(res, com)
	}
	if len(res) > 0 {
		resp.Items = res
	}

	resp.Base = &common.BaseResponse{
		Code: e.SuccessCode,
		Msg:  e.SuccessMsg,
	}

	return resp, nil
}

func DeleteComment(c context.Context, ctx *app.RequestContext, commentReq *comment.CommentDeleteReq) (interface{}, error) {
	vid, _ := strconv.ParseInt(commentReq.GetVid(), 10, 64)
	cid, _ := strconv.ParseInt(commentReq.GetCid(), 10, 64)
	claim, err := utils.ParseToken(string(ctx.GetHeader("access_token")))
	if err != nil {
		return response.BadResponse(), err
	}
	if len(commentReq.GetVid()) != 0 {
		err = dao.Db.Model(&model.Comment{}).Where("uid = ?", claim.Uid).
			Where("vid = ?", vid).Delete(&model.Comment{}).Error
		if err != nil {
			return response.BadResponse(), err
		}
		err = dao.Db.Model(&model.Video{}).Where("vid = ?", vid).Update("comment_count", 0).Error
		if err != nil {
			return response.BadResponse(), err
		}
	} else if len(commentReq.GetCid()) != 0 {
		var c model.Comment
		err = dao.Db.Model(&model.Comment{}).Where("uid = ?", claim.Uid).
			Where("cid = ?", cid).First(&c).Delete(&model.Comment{}).Error
		if err != nil {
			return response.BadResponse(), err
		}
		err = dao.Db.Model(&model.Video{}).Where("vid = ?", c.Vid).UpdateColumn("comment_count", gorm.Expr("comment_count+?", -1)).Error
	}
	if len(commentReq.GetVid()) == 0 && len(commentReq.GetCid()) == 0 {
		return response.BadResponse(), errors.New("请重新操作")
	}

	return response.GoodResponse(), nil
}
