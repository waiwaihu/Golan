package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBL *gorm.DB

const (
	Username = "root"      //账号
	Password = "123456"    //密码
	Host     = "127.0.0.1" //数据库地址，可以是Ip或者域名
	Port     = 3306        //数据库端口
	Dbname   = "go_db"     //数据库名
	Timeout  = "10s"       //连接超时，10秒
)

type Studen struct {
	Name      string
	Age       int
	MyStudent string
}

func init() {
	//gorm的默认日志是只打印错误和慢SQL 我们可以自己设置
	var mysqlLogger logger.Interface
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", Username, Password, Host, Port, Dbname, Timeout)

	// 要显示的日志等级
	mysqlLogger = logger.Default.LogMode(logger.Info)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	fmt.Println(db)
	DBL = db
}

func main() {
	DBL.AutoMigrate(&Studen{})
}
