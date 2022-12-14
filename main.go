package main

import (
	config "JoinUs/configs"
	"JoinUs/dao"
	"JoinUs/rouers"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

func main() {
	InitConf()
	err := dao.InitMysql()
	if err != nil {
		return
	}

	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			panic(err)
		}
	}(dao.DB)

	InitGin()
	//热加载配置文件
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
func InitGin() {
	r := rouers.SetRouter()
	RunPort := viper.GetString("server.port")
	if RunPort != "" {
		err := r.Run(":" + RunPort)
		if err != nil {
			return
		}
	}
	err := r.Run()
	if err != nil {
		return
	}
}

func InitConf() {
	config.InitLogger()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./configs/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}
