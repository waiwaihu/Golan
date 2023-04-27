package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

// https://gitea.com/xorm/xorm/src/branch/master/README_CN.md
var (
	usernamel  string = "root"
	passwordl  string = "123456"
	ipAddressl string = "127.0.0.1"
	portl      int    = 3306
	dbNamel    string = "go_db"
	charsetl   string = "utf8mb4"
)

type Userl struct {
	Id      int64
	Name    string
	Age     int
	Avatar  string
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	//数据库基本信息
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", usernamel, passwordl, ipAddressl, portl, dbNamel, charsetl)
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println("数据库连接失败!")
	}
	err = engine.Sync(new(Userl)) //同步结构体，映射创建一个数据库
	if err != nil {
		fmt.Println("数据库同步失败!")
	}

	//Insert方法的参数可以是一个或多个Struct的指针，一个或多个Struct的Slice的指针
	// 插入一条或者多条记录
	user1 := Userl{Id: 1, Name: "tom", Age: 20, Avatar: "猫", Passwd: "123456", Created: time.Time{}, Updated: time.Time{}}
	user2 := Userl{Id: 2, Name: "jerry", Age: 18, Avatar: "老鼠", Passwd: "123456", Created: time.Time{}, Updated: time.Time{}}
	n, err := engine.Insert(&user1)
	//n, err := engine.Insert(&user1,&user2)

	//插入切片
	var user3 []Userl
	user3 = append(user3, user1)
	user3 = append(user3, user2)
	//n, err := engine.Insert(&user3)

	//engine.Insert插入对象，返回值：受影响的行数
	if err != nil {
		fmt.Println("数据插入失败")
	}
	if n > 1 {
		fmt.Println("数据插入成功")
	}
}
