package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DBS *gorm.DB
var newLogger logger.Interface

type Stud struct {
	Name      string
	Age       int
	MyStudent string
}

func init() {
	username := "root"   //账号
	password := "123456" //密码
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "go_db"    //数据库名
	timeout := "10s"     //连接超时，10秒

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	// 要显示的日志等级
	newLogger = logger.Default.LogMode(logger.Info)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	fmt.Println(db)
	DBS = db
}

func main() {
	var model Stud
	session := DBS.Session(&gorm.Session{Logger: newLogger})

	DBS.AutoMigrate(&Stud{})

	session.First(&model)
	// SELECT * FROM `studs` ORDER BY `studs`.`name` LIMIT 1

	//如果只想某些语句显示日志
	//DBS.Debug().AutoMigrate(&Stud{})
	//DBS.Debug().First(&model)
}
