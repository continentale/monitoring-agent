package types

type Procs struct {
	Name          string   `json:"name"`
	MemoryPercent float32  `json:"memoryPercent"`
	CPUPercent    float64  `json:"cpuPercent"`
	Exe           string   `json:"exe"`
	Status        []string `json:"status"`
}
