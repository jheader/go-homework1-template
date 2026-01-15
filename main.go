package main

import (
	"fmt"

	"github.com/jheader/go-homework1-template/util"
	"gorm.io/gorm"
)

// 定义测试模型
type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {

	// 替换为你的 MySQL 连接信息（格式：用户名:密码@tcp(地址:端口)/数据库名?charset=utf8mb4&parseTime=True&loc=Local）
	//dsn := "root:123456@tcp(192.168.101.36:3306)/test_db?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	db, err := util.NewUtilDB()
	if err != nil {
		panic("连接数据库失败: " + err.Error())
	}
	// 自动迁移表结构
	err = db.AutoMigrate(&User{})
	if err != nil {
		panic("迁移表失败: " + err.Error())
	}

	fmt.Println("GORM 安装并连接数据库成功！")
}
