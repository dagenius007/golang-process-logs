package types

import "time"

type Process struct {
	ID                    int       `json:"id"`
	User                  string    `json:"user"`
	PID                   int32     `json:"pid"`
	CpuUsage              float32   `json:"cpuUsage"`
	MemoryPercentageUsage float32   `json:"memoryPercentageUsage"`
	VirtualMemorySize     int64     `json:"virtualMemorySize"`
	ResidentMemorySize    int64     `json:"residentMemorySize"`
	Tty                   string    `json:"tty"`
	State                 string    `json:"state"`
	Started               string    `json:"started"`
	TotalTime             string    `json:"totalTime"`
	Command               string    `json:"command"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
}
