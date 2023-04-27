// 更新数据，默认只更新非空和非0 的字段
// 删除记录，需要注意，删除必须至少有一个条件，否则会报错
// exec执行一个SQL语句
package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

var (
	usernamee  string = "root"
	passworde  string = "123456"
	ipAddresse string = "127.0.0.1"
	porte      int    = 3306
	dbNamee    string = "go_db"
	charsete   string = "utf8mb4"
)

type Usere struct {
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
	err = engine.Sync(new(Usere)) //同步结构体，映射创建一个数据库
	if err != nil {
		fmt.Println("数据库同步失败!")
	}

	user := Usere{Name: "tom", Age: 25}
	//（1）更新id为1的数据
	n, err := engine.ID(1).Update(&user)
	//n, _ := engine.Update(&user, &User{Passwd: "123456"})
	fmt.Println(n)

	//（2）删除id为1的数据
	a, err := engine.ID(1).Delete(&Userl{})
	if err != nil {
		fmt.Println("数据更新失败")
	}
	fmt.Println(a)

	//(3)Exec 执行一个SQL语句
	engine.Exec("update user set age = ? where name = ?", 24, "tomm")
}
