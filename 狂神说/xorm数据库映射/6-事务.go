package main

import (
	"fmt"
	"xorm.io/xorm"
)

func mai() error {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4")
	if err != nil {
		fmt.Printf("xorm数据库连接失败：%v\n", err)
		return err
	}

	//在一个Go程中有事务
	session := engine.NewSession()
	defer session.Close()

	// add Begin() before any action
	if err := session.Begin(); err != nil {
		// if returned then will rollback automatically
		return err
	}

	user1 := Userl{}
	if _, err := session.Insert(&user1); err != nil {
		return err
	}

	user2 := Userl{}
	if _, err := session.Where("id = ?", 2).Update(&user2); err != nil {
		return err
	}

	if _, err := session.Exec("delete from userinfo where username = ?", user2.Name); err != nil {
		return err
	}

	// add Commit() after all actions
	return session.Commit()

}
