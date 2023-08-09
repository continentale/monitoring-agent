package utils

import "github.com/continentale/monitoring-agent/config"

func GetRightExec(search string, commands []config.ExecEntries) *config.ExecEntries {
	for _, value := range commands {
		if value.Name == search {
			return &value
		}
	}
	return nil
}

func GetRightFile(search string, commands []config.FileEntries) *config.FileEntries {
	for _, value := range commands {
		if value.Name == search {
			return &value
		}
	}
	return nil
}
