package paths

import (
	"time"

	"github.com/spf13/viper"
)

func InitCommon() {
	viper.SetDefault("server.address", "")
	viper.SetDefault("server.port", 10240)
	viper.SetDefault("server.protocol", "http")
	viper.SetDefault("server.certificate", "certificates/cert.crt")
	viper.SetDefault("server.key", "certificates/priv.key")

	viper.SetDefault("server.secret", "superSecretPassword")
	viper.SetDefault("server.useSecret", false)

	// application specific
	viper.SetDefault("global.timeStringFormat", time.UnixDate)

	// endpoint specific

	// mem
	viper.SetDefault("mem.enabled", true)

	// procs
	viper.SetDefault("procs.enabled", true)

	// disk
	viper.SetDefault("disk.enabled", true)

	// load
	viper.SetDefault("load.enabled", true)

	// time
	viper.SetDefault("time.enabled", true)

	// cpu
	viper.SetDefault("cpu.enabled", true)
	viper.SetDefault("cpu.perCPU", true)

	// file
	viper.SetDefault("file.enabled", true)

	// exec
	viper.SetDefault("exec.enabled", true)

}
