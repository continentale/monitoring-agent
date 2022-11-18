package types

import "github.com/shirou/gopsutil/v3/cpu"

type CPUS struct {
	TimesStat cpu.TimesStat `json:"timesStat"`
	Usage     float64       `json:"usage"`
}
