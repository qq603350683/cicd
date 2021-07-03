package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

var DB *gorm.DB

func ConnectMysql() error {
	// test
	// connString 用户名:密码@(主机地址:端口)/数据库名称?charset=utf8&parseTime=True&loc=Local
	connString := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", "root", "root", "127.0.0.1", 3306, "test_micro")
	log.Println(connString)
	db, err := gorm.Open("mysql", connString)
	if err != nil {
		return err
	}

	// 禁止复表
	db.SingularTable(true)

	// 设置连接池
	// 空闲
	db.DB().SetMaxIdleConns(20)

	// 打开
	db.DB().SetMaxOpenConns(20)

	// 超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	// 开启SQL打印模式
	db.LogMode(true)

	DB = db

	migration()

	return nil
}

func migration() {
	set := "ENGINE=InnoDB"

	user := NewUserTest()

	if DB.HasTable(user) {
		DB.AutoMigrate(user)
	} else {
		DB.Set("gorm:table_options", set).CreateTable(user)
	}
}