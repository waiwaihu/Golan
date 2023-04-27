package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DBA *gorm.DB

type Stude struct {
	Name      string
	Age       int
	MyStudent string
}

// 如果你想自定义日志的显示
func init() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // （日志输出的目标，前缀和日志包含的内容）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 使用彩色打印
		},
	)
	username := "root"   //账号
	password := "123456" //密码
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "go_db"    //数据库名
	timeout := "10s"     //连接超时，10秒

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	fmt.Println(db)
	DBA = db
}
func main() {
	DBA.AutoMigrate(&Stude{})
}
