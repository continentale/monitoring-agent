package types

import (
	"github.com/shirou/gopsutil/v3/disk"
)

type Disks struct {
	Usage   *disk.UsageStat    `json:"usage"`
	Details disk.PartitionStat `json:"details"`
}
