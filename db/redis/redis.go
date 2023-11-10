package redis

import (
	"fmt"
	"strconv"
	"vi-server/utils"

	"github.com/gomodule/redigo/redis"
)

var (
	conn redis.Conn
)

func InitRedis() error {
	var err error
	config := utils.GetConfig()
	conn, err = redis.Dial("tcp", config.Db.Redis.Host+":"+strconv.Itoa(config.Db.Redis.Port))
	if err != nil {
		fmt.Println("Connect to Redis error", err)
		return err
	}
	defer conn.Close()

	// 测试连接
	pong, err := redis.String(conn.Do("PING"))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(pong)
	return nil
}

func Exec(commandName string, args ...interface{}) (interface{}, error) {
	result, err := conn.Do(commandName, args...)
	if err != nil {
		return nil, err
	}
	return result, nil
}
