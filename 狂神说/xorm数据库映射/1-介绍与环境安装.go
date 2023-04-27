package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

//官网：https://xorm.io/
//go get xorm.io/xorm
//https://gitea.com/xorm/xorm/src/branch/master/README_CN.md

/*
DROP TABLE IF EXISTS user;
CREATE TABLE user (

	id int(11) NOT NULL AUTO_INCREMENT COMMENT '自增主键',
	userid int(11) NULL DEFAULT NULL COMMENT '用户id',
	username VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名',
	password VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户密码',
	avatar VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户头像',
	create_time datetime(0) NULL DEFAULT NULL COMMENT '创建时间',
	update_time datetime(0) NULL DEFAULT NULL COMMENT '更新时间',
	PRIMARY KEY (id) USING BTREE

)ENGINE=INNODB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci  ROW_FORMAT=Dynamic;
*/
var (
	username  string = "root"
	password  string = "123456"
	ipAddress string = "127.0.0.1"
	port      int    = 3306
	dbName    string = "go_db"
	charset   string = "utf8mb4"
)

type User struct {
	Id      int64
	Name    string
	Age     int
	Avatar  string
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	//构建数据库信息
	//"root:123456@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4"
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", username, password, ipAddress, port, dbName, charset)
	engine, err := xorm.NewEngine("mysql", dataSourceName)
	if err != nil {
		fmt.Println("数据库连接失败!")
	}
	err = engine.Sync(new(User)) //同步结构体，映射创建一个数据库
	if err != nil {
		fmt.Println("数据库同步失败!")
	}
}
