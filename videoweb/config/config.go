package config

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	SqlUserName            string
	SqlPassword            string
	DataBase               string
	MysqlIP                string
	RabbitmqUserName       string
	RabbitmqPassword       string
	RabbitmqIP             string
	ExchangeName           string
	RedisAddr              string
	RedisPassword          string
	RedisDB                int
	EsAddr                 string
	EsName                 string
	EsPassword             string
	CertificateFingerprint string //第一次打开es时可获得
	EndPoint               string
	AccessKeyID            string
	SecretAccessKey        string
	SSL                    bool
	BucketName             string
)

func InitMysql(InitFile *ini.File) {

	SqlUserName = InitFile.Section("mysql").Key("SqlUserName").String()
	SqlPassword = InitFile.Section("mysql").Key("SqlPassword").String()
	DataBase = InitFile.Section("mysql").Key("DataBase").String()
	MysqlIP = InitFile.Section("mysql").Key("MysqlIP").String()
}

func InitRabbitMQ(InitFile *ini.File) {
	RabbitmqUserName = InitFile.Section("RabbitMQ").Key("RabbitmqUserName").String()
	RabbitmqPassword = InitFile.Section("RabbitMQ").Key("RabbitmqPassword").String()
	RabbitmqIP = InitFile.Section("RabbitMQ").Key("RabbitmqIP").String()
	ExchangeName = InitFile.Section("RabbitMQ").Key("ExchangeName").String()
}

func InitRedis(InitFile *ini.File) {
	RedisAddr = InitFile.Section("redis").Key("RedisAddr").String()
	RedisPassword = InitFile.Section("redis").Key("RedisPassword").String()
	RedisDB = InitFile.Section("redis").Key("RedisDB").MustInt()
}

func InitEs(InitFile *ini.File) {
	EsAddr = InitFile.Section("elasticsearch").Key("EsAddr").String()
	EsName = InitFile.Section("elasticsearch").Key("EsName").String()
	EsPassword = InitFile.Section("elasticsearch").Key("EsPassword").String()
	CertificateFingerprint = InitFile.Section("elasticsearch").Key("CertificateFingerprint").String()
}

func InitMinIo(InitFile *ini.File) {
	EndPoint = InitFile.Section("MinIo").Key("EndPoint").String()
	AccessKeyID = InitFile.Section("MinIo").Key("AccessKeyID").String()
	SecretAccessKey = InitFile.Section("MinIo").Key("SecretAccessKey").String()
	SSL = InitFile.Section("MinIo").Key("SSL").MustBool()
	BucketName = InitFile.Section("MinIo").Key("BucketName").String()
}

func Init() {

	InitFile, err := ini.Load("./config/config.ini")
	if err != nil {
		log.Panicln(err)
	}
	InitMysql(InitFile) //
	InitMinIo(InitFile) //
	InitRabbitMQ(InitFile)
	InitRedis(InitFile)
	InitEs(InitFile)
}
