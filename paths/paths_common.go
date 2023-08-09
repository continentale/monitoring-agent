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
	viper.SetDefault("endpoints.mem.enabled", true)

	// procs
	viper.SetDefault("endpoints.procs.enabled", true)
	viper.SetDefault("endpoints.procs.filter", ".")

	// disks
	viper.SetDefault("endpoints.disks.enabled", true)
	viper.SetDefault("endpoints.disks.filter", ".")

	// load
	viper.SetDefault("endpoints.load.enabled", true)

	// time
	viper.SetDefault("endpoints.time.enabled", true)

	// cpus
	viper.SetDefault("endpoints.cpus.enabled", true)
	viper.SetDefault("endpoints.cpus.perCPU", true)
	viper.SetDefault("endpoints.cpus.filter", ".")

	// file
	viper.SetDefault("endpoints.file.enabled", false)

	// exec
	viper.SetDefault("endpoints.exec.enabled", false)

	// version
	viper.SetDefault("endpoints.version.enabled", true)
}
