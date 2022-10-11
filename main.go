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

	viper.WatchConfig()

	r := gin.Default()

	v2 := r.Group("/api/v2")

	v2.Use(api.AuthorizationV2)

	v2.GET("mem", api.GetMemory)
	v2.GET("procs", api.GetProcs)
	v2.GET("disk", api.GetDisk)
	v2.GET("load", api.GetLoad)
	v2.GET("time", api.GetTime)
	v2.GET("cpu", api.GetCPU)

	// TODO implement this functions
	// v2.GET("file", api.ShowFile)
	// v2.GET("exec", api.ExecCommand)

	if viper.GetString("server.protocol") == "https" {
		r.RunTLS(fmt.Sprintf("%s:%d", viper.GetString("server.address"), viper.GetInt("server.port")), viper.GetString("server.certificate"), viper.GetString("server.key"))
	} else {
		r.Run(fmt.Sprintf("%s:%d", viper.GetString("server.address"), viper.GetInt("server.port")))
	}
}
