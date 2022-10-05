package paths

import "github.com/spf13/viper"

func InitCommon() {
	viper.SetDefault("address", "*")
	viper.SetDefault("port", 10240)
}
