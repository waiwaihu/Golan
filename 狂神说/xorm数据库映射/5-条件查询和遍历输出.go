package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

type UserA struct {
	Id      int64
	Name    string
	Age     int
	Avatar  string
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4")
	if err != nil {
		fmt.Printf("xorm数据库连接失败：%v\n", err)
		return
	} else {
		err := engine.Ping()
		if err != nil {
			fmt.Printf("连接成功，ping失败:%v\n", err)
		} else {
			fmt.Println("连接成功，ping通了...")
		}
	}

	//(1)Query 最原始的也支持SQL语句查询，返回的结果类型为 []map[string][]byte。
	//QueryString 返回 []map[string]string, QueryInterface 返回 []map[string]interface{}
	r, _ := engine.QueryString("select * from user")
	fmt.Println(r)

	//(2)get查询单条记录
	user := User{Name: "tom"}
	engine.Where("name = ?", user.Name).Desc("id").Get(&user)
	// SELECT * FROM user WHERE name = ? ORDER BY id DESC LIMIT 1
	fmt.Println(user)

	var name string
	userl := Userl{}
	engine.Table(&userl).Where("id = ?", 1).Cols("name").Get(&name)
	// SELECT name FROM user WHERE id = ?
	fmt.Println(name)

	//(3)Find 查询多条记录，当然可以使用Join和extends来组合使用
	var users []User
	engine.Where("name = ?", name).And("age > 10").Limit(10, 0).Find(&users)
	// SELECT * FROM user WHERE name = ? AND age > 10 limit 10 offset 0
	fmt.Println(users)

	//(4)Count 获取记录条数
	counts, _ := engine.Count(&user)
	// SELECT count(*) AS total FROM user
	fmt.Println(counts)

	//(5)Iterate 和 Rows 根据条件遍历数据库，可以有两种方式: Iterate and Rows
	engine.Iterate(&User{Name: name}, func(idx int, bean interface{}) error {
		user := bean.(*User)
		fmt.Println(user) //循环遍历
		return nil
	})
	// SELECT * FROM user

	rows, _ := engine.Rows(&User{Name: name})
	// SELECT * FROM user
	defer rows.Close()
	bean := new(Userl)
	for rows.Next() {
		//err = rows.Scan(bean)
		rows.Scan(bean)
		fmt.Println(bean)
	}

}
