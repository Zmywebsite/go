package utils

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	App *App // 配置程序信息
	Db  *Db  // 配置数据库信息
}
type App struct {
	Port  int    `json:"port"`  // 程序端口号
	Token string `json:"token"` // 请求token
}
type Db struct {
	Redis *Redis // 配置redis信息
}
type Redis struct {
	Mode string `json:"mode"` // 集群模式
	Host string `json:"host"` // 地址
	Port int    `json:"port"` // 端口号
}

var config Config

// 初始化配置文件参数
func InitConfig() {
	// 初始化 viper 配置// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// 设置配置文件的路径为config.yaml的绝对路径
	configPath := filepath.Join(wd, "config.yaml")
	viper.SetConfigFile(configPath)

	// 读取配置文件
	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 将配置文件映射到结构体中
	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}

// 获取参数
func GetConfig() Config {
	return config
}

func SetConfig() {

}
