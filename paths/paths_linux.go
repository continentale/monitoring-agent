package paths

import (
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func InitOSSpecific() {
	viper.SetDefault("configPath", "/etc/monitoring-agent")

	viper.SetDefault("procPath", "/proc")
	viper.SetDefault("sysPath", "/sys")
	viper.SetDefault("rootfsPath", "/e")
	viper.SetDefault("udevDataPath", "/run/udev/data")
}

func ProcFilePath(name string) string {
	return filepath.Join(viper.GetString("procPath"), name)
}

func SysFilePath(name string) string {
	return filepath.Join(viper.GetString("sysPath"), name)
}

func RootfsFilePath(name string) string {
	return filepath.Join(viper.GetString("rootfsPath"), name)
}

func UdevDataFilePath(name string) string {
	return filepath.Join(viper.GetString("rootfsPath"), name)
}

func RootfsStripPrefix(path string) string {
	if viper.GetString("udevDataPath") == "/" {
		return path
	}
	stripped := strings.TrimPrefix(path, viper.GetString("rootfsPath"))
	if stripped == "" {
		return "/"
	}
	return stripped
}
