package main

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func initConfig() {
	prepareDefaultPaths()

	viper.SetConfigName("config")
	viper.SetConfigType("toml")

}

func prepareDefaultPaths() {
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")
	viper.AddConfigPath(viper.GetString("configPath"))

	// get home dir for config path location
	dirname, err := os.UserHomeDir()
	if err != nil {
		log.Printf("ERROR %v", err)
	}

	viper.AddConfigPath(dirname)
}
