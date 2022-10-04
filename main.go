package main

import (
	"github.com/spf13/viper"
)

func main() {
	initConfig()
	viper.ReadInConfig()

	/*
		r := gin.Default()

		r.GET("/api/ram", getRAM)

		r.Run()
	*/
}
