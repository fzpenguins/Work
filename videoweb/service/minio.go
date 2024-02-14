package service

import (
	"context"
	"errors"
	"fmt"
	"path/filepath"
	"time"
	"videoweb/config"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v7"
)

func GetURL(objectName string) (string, error) {
	ctx := context.Background()
	_, err := config.MinioClient.BucketExists(ctx, config.BucketName)
	if err != nil {
		return "", errors.New("minio配置错误")
	}
	presignedURL, err := config.MinioClient.PresignedGetObject(ctx, config.BucketName,
		objectName, time.Second*300, nil)
	if err != nil {
		return "", errors.New("minio配置错误")
	}
	return presignedURL.String(), nil

}

func UploadVideo(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	objectName := "video/" + uuid.Must(uuid.NewRandom()).String() + ext
	_, err := config.MinioClient.FPutObject(context.Background(), config.BucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		return "", errors.New("文件上传失败")
	}
	return objectName, nil
}

func UploadCover(filePath string) (string, error) {
	ext := filepath.Ext(filePath)
	objectName := "cover/" + uuid.Must(uuid.NewRandom()).String() + ext
	_, err := config.MinioClient.FPutObject(context.Background(), config.BucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		return "", errors.New("文件上传失败")
	}
	return objectName, nil
}

func GenerateVideoName(userID int64) string {
	currentTime := time.Now()
	// 获取年月日和小时分钟
	year, month, day := currentTime.Date()
	hour, minute := currentTime.Hour(), currentTime.Minute()
	return fmt.Sprintf("%v_%d%02d%02d_%02d%02d_video.mp4", userID, year, month, day, hour, minute)
}
func GenerateCoverName(userID int64) string {
	currentTime := time.Now()
	// 获取年月日和小时分钟
	year, month, day := currentTime.Date()
	hour, minute := currentTime.Hour(), currentTime.Minute()
	return fmt.Sprintf("%v_%d%02d%02d_%02d%02d_cover.jpg", userID, year, month, day, hour, minute)
}
