package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
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

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, mem)
}

func GetProcs(c *gin.Context) {
	procs, _ := process.Processes()

	results := make([]types.Procs, len(procs))

	for i, p := range procs {
		results[i].Name, _ = p.Name()
		results[i].MemoryPercent, _ = p.MemoryPercent()
		results[i].Exe, _ = p.Exe()
		results[i].CPUPercent, _ = p.CPUPercent()
		results[i].Status, _ = p.Status()
	}

	jsonData, _ := json.Marshal(results)

	filter, ok := c.GetQueryArray("filter")
	if !ok {
		filter = append(filter, viper.GetString("procs.filter"))
	}

	if len(filter) > 1 || filter[0] != "." { // . is the default filter which means no filter at all
		jsonData = types.ApplyFilter(jsonData, filter, "procs")
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}

func GetDisk(c *gin.Context) {
	disks, _ := disk.Partitions(true)

	results := make([]types.Disks, len(disks))

	for i, d := range disks {
		results[i].Details = d
		results[i].Usage, _ = disk.Usage(d.Mountpoint)
	}

	jsonData, _ := json.Marshal(results)

	filter, ok := c.GetQueryArray("filter")
	if !ok {
		filter = append(filter, viper.GetString("disks.filter"))
	}

	if len(filter) > 1 || filter[0] != "." { // . is the default filter which means no filter at all
		jsonData = types.ApplyFilter(jsonData, filter, "disks")
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}

func GetLoad(c *gin.Context) {
	load, _ := load.Avg()

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, load)
}

func GetTime(c *gin.Context) {
	var result types.TimeSync

	timeNow := time.Now()

	result.Timestamp = timeNow.Unix()
	result.Formatted = timeNow.Format(viper.GetString("global.timeStringFormat"))
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, result)
}

func GetCPU(c *gin.Context) {
	perCPU, _ := strconv.ParseBool(c.DefaultQuery("perCPU", viper.GetString("cpus.perCPU")))
	cpus, _ := cpu.Times(perCPU)

	usagesPercent, _ := cpu.Percent(time.Second, perCPU)

	result := make([]types.CPUS, len(cpus))

	for i := range result {
		result[i].TimesStat = cpus[i]
		result[i].Usage = usagesPercent[i]
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, result)
}

func ShowFile(c *gin.Context) {
	name := c.DefaultQuery("name", "")

	if name == "" {
		c.AbortWithError(http.StatusBadRequest, errors.New("name not defined"))
	}

	filePath := viper.GetString("file.entries." + name)

	if filePath == "" {

		filePath = viper.GetString("file.entries." + name + ".path")

		if filePath == "" {
			c.AbortWithError(http.StatusNotFound, errors.New("command not defined"))
		}
	}

	if viper.GetBool("file.entries." + name + ".contentOnly") {
		fileContent, _ := ioutil.ReadFile(filePath)
		c.Data(http.StatusOK, "text/html; charset=UTF-8", fileContent)
	} else {
		file, _ := os.OpenFile(filePath, os.O_RDONLY, 0644)
		stat, _ := file.Stat()

		fileContent, _ := ioutil.ReadFile(filePath)

		customFile := types.File{
			Path:    filePath,
			IsDir:   stat.IsDir(),
			ModTime: stat.ModTime().Unix(),
			Mode:    stat.Mode().String(),
			Name:    stat.Name(),
			Size:    stat.Size(),
			Content: string(fileContent),
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, customFile)
	}
}

func ExecCommand(c *gin.Context) {
	var waitStatus syscall.WaitStatus
	var check types.Check

	name := c.DefaultQuery("name", "")

	if name == "" {
		c.AbortWithError(http.StatusNotFound, errors.New("name not defined"))
		return
	}

	commandPath := viper.GetString("exec.entries." + name)

	if commandPath == "" {
		commandPath = viper.GetString("exec.entries." + name + ".path")
		if commandPath == "" {
			c.AbortWithError(http.StatusBadRequest, errors.New("command not defined"))
			return
		}
	}

	path := viper.GetString("exec.entries." + name + ".shell")
	if path == "" {
		// the default is not overwritten
		path = viper.GetString("exec.shell")
	}

	cmd := exec.Command(path, strings.Fields(commandPath)...)

	out, err := cmd.CombinedOutput()

	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
		} else {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
		}
	} else {
		waitStatus = cmd.ProcessState.Sys().(syscall.WaitStatus)
	}

	outString := string(out)
	check.Output = outString
	check.ExitCode = waitStatus.ExitStatus()

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, check)
}
