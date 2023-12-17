package processes

import "os"

// . "binalyze-test/types"

type ProcessInfo struct {
	PID                   int32
	Ppid                  int32
	CpuUsage              float32
	MemoryPercentageUsage float32
}

// func statInfo(pid int32) (any, error) {
// 	fmt.Println("pid", pid)
// 	statPath := fmt.Sprintf("/proc/%d/stat", pid)

// 	dataBytes, err := ioutil.ReadFile(statPath)
// 	if err != nil {
// 		return nil, nil
// 	}

// 	fmt.Println("data", dataBytes)
// 	return nil, nil
// }

func GetProcesses() ([]any, error) {
	os.Open("/proc")
	return getProcessList()
}
