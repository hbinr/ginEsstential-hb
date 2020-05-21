//File  : main.go
//Author: duanhaobin
//Date  : 2020/5/20

package main

import (
	"ginEssential-hb/common"
	"ginEssential-hb/router"
	"os"

	"github.com/spf13/viper"

	"github.com/gin-gonic/gin"
)

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()

	// 强制在控制台中输出颜色
	r = router.CollectRoute(r)

	// 设置端口，如果读取不到，则默认端口
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	curretPath, _ := os.Getwd()
	// 设置配置文件名
	viper.SetConfigName("application")
	// 设置文件类型
	viper.SetConfigType("yml")
	// 设置目录
	viper.AddConfigPath(curretPath + "/config")
	// 读取文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}
