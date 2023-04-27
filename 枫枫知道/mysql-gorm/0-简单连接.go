package main

//http://docs.fengfengzhidao.com/#/docs/gorm%E6%96%87%E6%A1%A3/3.%E5%8D%95%E8%A1%A8%E6%9F%A5%E8%AF%A2
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	username := "root"   //账号
	password := "123456" //密码
	host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	port := 3306         //数据库端口
	Dbname := "go_db"    //数据库名
	timeout := "10s"     //连接超时，10秒

	// root:root@tcp(127.0.0.1:3306)/gorm?
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	fmt.Println(db)
}
