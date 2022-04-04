package cfg

import (
	"local/controllers"

	"github.com/spf13/viper"
)

type Config struct {
	Port, Width, Height int
}

func GetConfig() *Config {
	v := viper.New()
	// viper读config.ini文件
	v.AddConfigPath(controllers.Dir()) // 添加配置文件路径
	v.SetConfigName("config")          // 配置文件名
	v.SetConfigType("ini")             // 配置文件类型

	// 读取配置文件
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 当有节时[server],读取时加上节
	// Config.Port = v.GetInt("server.port")
	// Config.Width = v.GetInt("windows.width")
	// Config.Height = v.GetInt("windows.height")
	c := Config{
		Port:   v.GetInt("server.port"),
		Width:  v.GetInt("windows.width"),
		Height: v.GetInt("windows.height"),
	}
	return &c

}
