package main

import (
	"strconv"
	"vi-server/db/redis"
	"vi-server/mod"
	"vi-server/utils"
)

// @title Go-Web学习开发记录
// @version 0.0.1
// @description 学习GoLang Web开发记录
// @host         localhost:8080
func main() {

	// swag init 更新接口文档

	// 初始化读取配置文件
	utils.InitConfig()
	config := utils.GetConfig() // 获取配置文件
	router := mod.Setupr()      // 初始化mod
	_ = redis.InitRedis()

	router.Run(":" + strconv.Itoa(config.App.Port))
}
