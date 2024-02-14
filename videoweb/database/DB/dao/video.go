package dao

import (
	"context"
	"videoweb/database/DB/model"

	"gorm.io/gorm"
)

type VideoDao struct {
	*gorm.DB
}

func NewVideo(ctx context.Context) *VideoDao {
	if ctx == nil {
		ctx = context.Background()
	}
	return &VideoDao{NewDBClient(ctx)}
}

func (dao *VideoDao) CreateVideo(video *model.Video) (err error) {
	err = dao.DB.Model(&model.Video{}).Create(&video).Error
	return
}

func (dao *VideoDao) FindVideoByVid(vid int64) (video model.Video, err error) {
	//err = dao.DB.Where("vid=?", vid).First(video).Error
	err = dao.DB.Model(&model.Video{}).Where("vid = ?", vid).First(&video).Error
	return
}
