package dao

import (
	"context"
	"fmt"
	"videoweb/database/DB/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func InitDB() {
	//config.Init()

	//dsn := strings.Join([]string{config.SqlUserName, ":", config.SqlPassword,
	//	"@tcp(" + config.MysqlIP + ")/", config.DataBase, //"")
	//	"?charset=utf8mb4&parseTime=True&loc=Local"}, "")
	const (
		SqlUserName = "root"
		SqlPassword = "123456"
		DataBase    = "videoWebsite"
		MysqlIP     = "192.168.1.105:3306" //"[fe80::7f4d:3db3:fe1:a2e1%16]:3306" // 使用容器名称而不是本地主机地址
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", SqlUserName, SqlPassword, MysqlIP, DataBase)

	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Migrate()
}

func Migrate() {
	Db.AutoMigrate(&model.User{})
	Db.AutoMigrate(&model.Video{})
	Db.AutoMigrate(&model.Relation{})
	Db.AutoMigrate(&model.Comment{})
	Db.AutoMigrate(&model.Like{})
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := Db
	return db.WithContext(ctx)
}
