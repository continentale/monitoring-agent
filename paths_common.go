package main

import "github.com/spf13/viper"

func init() {
	viper.SetDefault("address", "*")
	viper.SetDefault("port", 10240)
}
