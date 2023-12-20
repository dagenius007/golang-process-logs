package types

import "time"

type Process struct {
	ID                 int       `json:"id"`
	User               string    `json:"user"`
	PID                int32     `json:"pid"`
	CpuUsage           float64   `json:"cpuUsage"`
	MemoryUsage        float64   `json:"memoryUsage"`
	ResidentMemorySize int64     `json:"residentMemorySize"`
	VirtualMemorySize  int64     `json:"virtualMemorySize"`
	State              string    `json:"state"`
	TotalTime          string    `json:"totalTime"`
	CpuTime            string    `json:"cpuTime"`
	Command            string    `json:"command"`
	Priority           string    `json:"priority"`
	CreatedAt          time.Time `json:"createdAt"`
	UpdatedAt          time.Time `json:"updatedAt"`
}

type ProcessUserReport struct {
	User                 string  `json:"user"`
	TotalUserCpuUsage    float64 `json:"totalUserCpuUsage"`
	TotalUserMemoryUsage float64 `json:"totalUserMemoryUsage"`
	TotalProcesses       int64   `json:"totalProcesses"`
}
