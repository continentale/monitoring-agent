package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/continentale/monitoring-agent/api"
	"github.com/continentale/monitoring-agent/paths"
)

func main() {
	initConfig()
	paths.InitCommon()
	paths.InitOSSpecific()
	viper.ReadInConfig()

	fmt.Println("sysctl", viper.GetString("sysPath"))

	r := gin.Default()
	r.GET("/api/mem", api.GetMemory)

	r.Run("127.0.0.1:10240")
}
