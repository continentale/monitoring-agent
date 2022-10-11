package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/continentale/monitoring-agent/types"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

func GetMemory(c *gin.Context) {
	mem, _ := mem.VirtualMemory()

	c.JSON(http.StatusOK, mem)
}

func GetProcs(c *gin.Context) {
	procs, _ := process.Processes()

	result := make([]types.Procs, len(procs))

	for i, p := range procs {
		result[i].Name, _ = p.Name()
		result[i].MemoryPercent, _ = p.MemoryPercent()
		result[i].Exe, _ = p.Exe()
		result[i].CPUPercent, _ = p.CPUPercent()
		result[i].Status, _ = p.Status()
	}

	c.JSON(http.StatusOK, result)
}

func GetDisk(c *gin.Context) {
	disks, _ := disk.Partitions(true)

	results := make([]types.Disks, len(disks))

	for i, d := range disks {
		results[i].Details = d
		results[i].Usage, _ = disk.Usage(d.Mountpoint)
	}
	c.JSON(http.StatusOK, results)
}

func GetLoad(c *gin.Context) {
	load, _ := load.Avg()

	c.JSON(http.StatusOK, load)
}

func GetTime(c *gin.Context) {
	var result types.TimeSync

	timeNow := time.Now()

	result.Timestamp = timeNow.Unix()
	result.Formatted = timeNow.Format(viper.GetString("timeStringFormat"))
	c.JSON(http.StatusOK, result)
}

func GetCPU(c *gin.Context) {
	perCPU, _ := strconv.ParseBool(c.DefaultQuery("perCPU", viper.GetString("cpu.perCPU")))
	cpus, _ := cpu.Times(perCPU)
	c.JSON(http.StatusOK, cpus)
}

func ShowFile(c *gin.Context) {
	load, _ := load.Avg()

	c.JSON(http.StatusOK, load)
}

func ExecCommand(c *gin.Context) {

}
