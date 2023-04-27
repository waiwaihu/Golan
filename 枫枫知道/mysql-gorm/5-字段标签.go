package main

type Stua struct {
	ID    uint // 默认使用ID作为主键
	Name  string
	Email *string // 使用指针是为了存空值
}

// 可以放多个 DB.AutoMigrate(&Student{}),默认的数据库类型太大了
//CREATE TABLE `f_students` (`id` bigint unsigned AUTO_INCREMENT,`name` longtext,`email` longtext,PRIMARY KEY (`id`))

/*
字段标签
type 定义字段类型

size 定义字段大小

column 自定义列名

primaryKey 将列定义为主键

unique 将列定义为唯一键

default 定义列的默认值

not null 不可为空

embedded 嵌套字段

embeddedPrefix 嵌套字段前缀

comment 注释

多个标签之前用 ; 连接
*/
type StudentInfo struct {
	Email  *string `gorm:"size:32"` // 使用指针是为了存空值
	Addr   string  `gorm:"column:y_addr;size:16"`
	Gender bool    `gorm:"default:true"`
}
type Studente struct {
	Name string      `gorm:"type:varchar(12);not null;comment:用户名"`
	UUID string      `gorm:"primaryKey;unique;comment:主键"`
	Info StudentInfo `gorm:"embedded;embeddedPrefix:s_"`
}

// 建表语句
/*CREATE TABLE `students` (
	`name` varchar(12) NOT NULL COMMENT '用户名',
	`uuid` varchar(191) UNIQUE COMMENT '主键',
	`s_email` varchar(32),
	`s_y_addr` varchar(16),
	`s_gender` boolean DEFAULT true,
	PRIMARY KEY (`uuid`)
)*/
