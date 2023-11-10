package user

import (
	"encoding/json"
	"fmt"
	"time"
	"vi-server/db/redis"
	_ "vi-server/docs"
)

// 用户登录
func userLogin(Name string, Password string) bool {
	if Name != "admin" {
		return false
	}

	if Password != "admin.Mgt1" {
		return false
	}

	return true
}

// 创建用户
func addUser(userName string, password string) bool {
	now := time.Now().Unix()
	userInfo := UserSchema{
		UserName:           userName,
		Password:           password,
		CreatTime:          int(now),
		UpdateTime:         int(now),
		LastLoginTime:      0,
		FirstLogin:         true,
		PasswordUpdateTime: int(now),
	}
	data, err := json.Marshal(userInfo)
	if err != nil {
		return false
	}
	result, err := redis.Exec("SET", userName, string(data))
	if err != nil {
		return false
	}
	fmt.Println(result)
	return true
}
