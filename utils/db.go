package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	//	// Db, err = gorm.Open(mysql.Open("root:L_1q2w3e4r@tcp(localhost)/datascs"))
	Db, err = gorm.Open(mysql.Open("root:password@tcp(localhost)/datascs"))
	//Db, err = gorm.Open(mysql.Open("root:12345678910@tcp(47.96.134.75)/datascs"))

	if err != nil {
		fmt.Println(err)
	}
	/*
			username := "root"       // 账号
		password := "L_1q2w3e4r" // 密码
		host := "8.130.52.233"   // 数据库地址，可以是Ip或者域名
		port := 3306             // 数据库端口
		Dbname := "datascs"      // 数据库名
		timeout := "10s"         // 连接超时，10秒

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
			username, password, host, port, Dbname, timeout)

		// 连接 Mysql, 获得 DB 类型实例，用于后面的数据库读写操作。
		Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("连接数据库失败, error=" + err.Error())
		}
	*/
}
