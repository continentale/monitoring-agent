package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Server    Server    `mapstructure:"server"`
	Global    Global    `mapstructure:"global"`
	Endpoints Endpoints `mapstructure:"endpoints"`
}
type Server struct {
	Protocol    string `mapstructure:"protocol"`
	Timeouts    int    `mapstructure:"timeouts"`
	Secret      string `mapstructure:"secret"`
	UseSecret   bool   `mapstructure:"useSecret"`
	Certificate string `mapstructure:"certificate"`
	Key         string `mapstructure:"key"`
	Port        int    `mapstructure:"port"`
	Address     string `mapstructure:"address"`
}
type Global struct {
	TimeStringFormat string `mapstructure:"timeStringFormat"`
}
type Version struct {
	Enabled bool `mapstructure:"enabled"`
}
type Mem struct {
	Enabled bool `mapstructure:"enabled"`
}
type Procs struct {
	Enabled bool `mapstructure:"enabled"`
}
type Disks struct {
	Enabled bool `mapstructure:"enabled"`
}
type Load struct {
	Enabled bool `mapstructure:"enabled"`
}
type Time struct {
	Enabled bool `mapstructure:"enabled"`
}
type Cpus struct {
	Enabled bool `mapstructure:"enabled"`
	PerCPU  bool `mapstructure:"perCPU"`
}
type FileEntries struct {
	Name        string `mapstructure:"name"`
	Path        string `mapstructure:"path"`
	ContentOnly bool   `yaml:"contentOnly,omitempty" mapstructure:"contentOnly"`
}
type File struct {
	Enabled bool          `mapstructure:"enabled"`
	Entries []FileEntries `mapstructure:"entries"`
}
type ExecEntries struct {
	Name  string `mapstructure:"name"`
	Path  string `mapstructure:"path"`
	Shell string `yaml:"shell,omitempty" mapstructure:"shell"`
}
type Exec struct {
	Enabled bool          `mapstructure:"enabled"`
	Shell   string        `mapstructure:"shell"`
	Path    string        `mapstructure:"path"`
	Entries []ExecEntries `mapstructure:"entries"`
}
type Endpoints struct {
	Version Version `mapstructure:"version"`
	Mem     Mem     `mapstructure:"mem"`
	Procs   Procs   `mapstructure:"procs"`
	Disks   Disks   `mapstructure:"disks"`
	Load    Load    `mapstructure:"load"`
	Time    Time    `mapstructure:"time"`
	Cpus    Cpus    `mapstructure:"cpus"`
	File    File    `mapstructure:"file"`
	Exec    Exec    `mapstructure:"exec"`
}

var (
	ConfigViper  *viper.Viper
	ConfigStruct *Config
)

func InitConfig() {
	PrepareDefaultPaths()

	viper.SetConfigName("monitoring-agent")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Cannot read config file", err)
	} else {
		log.Printf("used configfile: %s", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(&ConfigStruct); err != nil {
		log.Fatal(err)
	}
	viper.WatchConfig()

	ConfigViper = viper.GetViper()
}

func PrepareDefaultPaths() {
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
