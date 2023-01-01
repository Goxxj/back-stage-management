package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	// 连接mysql数据库，  gorm.Config有很多配置 例如前缀  后准 等等
	dsn := "root:qwe123@tcp(127.0.0.1:3306)/db1?charset=utf8&parseTime=True&loc=Local" // 数据库配置 root账号  0000密码  shop数据库  utf8mb4编码
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	// 数据库连接出错 退出
	if err != nil {
		fmt.Println(err)
	}

	DB.AutoMigrate(&Test{})

}
