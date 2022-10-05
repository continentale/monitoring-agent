package paths

import "github.com/spf13/viper"

func InitOSSpecific() {
	viper.SetDefault("configPath", "C:\\Program Files\\monitoring-agent")
}
