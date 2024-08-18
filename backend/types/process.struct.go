package types

import (
	"context"
	"time"

	"github.com/uptrace/bun"
)

type Process struct {
	bun.BaseModel      `bun:"table:processes"`
	Id                 int       `json:"id" bun:"id"`
	User               string    `json:"user" bun:"user"`
	PID                int32     `json:"pid" bun:"pid"`
	CpuUsage           float64   `json:"cpu_usage" bun:"cpu_usage"`
	MemoryUsage        float64   `json:"memory_usage" bun:"memory_usage"`
	ResidentMemorySize int64     `json:"resident_memory_size" bun:"resident_memory_size"`
	VirtualMemorySize  int64     `json:"virtual_memory_size" bun:"virtual_memory_size"`
	State              string    `json:"state" bun:"state"`
	TotalTime          string    `json:"total_time" bun:"total_time"`
	CpuTime            string    `json:"cpu_time" bun:"cpu_time"`
	Command            string    `json:"command" bun:"command"`
	Priority           string    `json:"priority" bun:"priority"`
	CreatedAt          time.Time `json:"created_at" bun:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" bun:"updated_at"`
}

type ProcessUserReport struct {
	User             string  `json:"user"`
	TotalCpuUsage    float64 `json:"totalCpuUsage"`
	TotalMemoryUsage float64 `json:"totalMemoryUsage"`
	TotalProcesses   int64   `json:"totalProcesses"`
}

type ProcessList struct {
	Processes []Process `json:"processes"`
	Totoal    int       `json:"total"`
	Limit     int       `json:"limit"`
	Page      int       `json:"page"`
}

type ProcessFilter struct {
	State  string
	User   string
	Search string
	Limit  int
	Offset int
}

type ProcessRepository interface {
	GetProcesses(ctx context.Context, filter ProcessFilter) ([]Process, int, error)
}
