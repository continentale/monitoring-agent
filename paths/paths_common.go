package paths

import (
	"time"

	"github.com/spf13/viper"
)

func InitCommon() {
	viper.SetDefault("server.address", "")
	viper.SetDefault("server.port", 20480)
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
	viper.SetDefault("procs.filter", ".")

	// disks
	viper.SetDefault("disks.enabled", true)
	viper.SetDefault("disks.filter", ".")

	// load
	viper.SetDefault("load.enabled", true)

	// time
	viper.SetDefault("time.enabled", true)

	// cpus
	viper.SetDefault("cpus.enabled", true)
	viper.SetDefault("cpus.perCPU", true)
	viper.SetDefault("cpus.filter", ".")

	// file
	viper.SetDefault("file.enabled", true)

	// exec
	viper.SetDefault("exec.enabled", true)

	// version
	viper.SetDefault("version.enabled", true)
}
