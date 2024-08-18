package process

import (
	"log"
	"strings"

	. "binalyze-test/types"
)

// This is not an actual CPU range
// This is a user defined priority and should not be used
func guagePriority(nice int) string {
	switch true {
	case nice >= -20 && nice <= -11:
		return "high"
	case nice >= -10 && nice <= 10:
		return "medium"
	default:
		return "low"
	}
}

func formatState(state string) string {
	switch true {
	case strings.Contains(state, "S") || strings.Contains(state, "D"):
		return "sleeping"
	case strings.Contains(state, "R"):
		return "running"
	case strings.Contains(state, "Z"):
		return "stopped_child"
	case strings.Contains(state, "T"):
		return "stopped"
	case strings.Contains(state, "I"):
		return "Idle"
	default:
		return "Unknown"
	}
}

/*
 This function gets all the current pids
 Get stats from running ps -o command with os/exec
 The values are mapped to the process struct accordily
*/

func getProcessList() []Process {
	processes := make([]Process, 0)
	pids, err := getPids()
	if err != nil {
		log.Println("Error fetching pids:", err)
		return processes
	}

	log.Println("Pids fetched successfully")

	for _, pid := range pids {
		p := &Process{}
		err = statInfo(p, pid)
		// Assumption is that an error here is caused by process not found
		// In that vain no process was outputted
		if err != nil {
			continue
		}

		processes = append(processes, *p)
	}

	log.Println("Processes fetched successfully")

	return processes
}

func GetProcesses() []Process {
	return getProcessList()
}
