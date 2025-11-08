package main

import (
	"fmt"
	"go-base-task-4/db"
	"go-base-task-4/models"
	"go-base-task-4/router"
)

func main() {
	//初始化数据库
	db.InitDB()
	defer func() {
		sqlDB, err := db.SqlDB.DB()
		if err == nil {
			sqlDB.Close()
		}
	}()

	//建表
	err := db.SqlDB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	if err != nil {
		fmt.Println("建表失败:", err)
		panic("建表失败：")
	}

	//启动服务
	router := router.InitRouter()
	router.Run(":8080")
}
