package dao

import (
	"context"
	"gorm.io/gorm"
	"videoweb/database/DB/model"
)

type CommentDao struct {
	*gorm.DB
}

func NewComment(ctx context.Context) *CommentDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &CommentDao{NewDBClient(ctx)}
}

func (dao *CommentDao) CreateComment(comment *model.Comment) (err error) {
	err = dao.DB.Model(&model.Comment{}).Create(&comment).Error
	return
}

func (dao *CommentDao) FindCommentByCid(cid int64) (comment model.Comment, err error) {
	err = dao.DB.Model(&model.Comment{}).Where("cid=?", cid).First(&comment).Error
	return
}
