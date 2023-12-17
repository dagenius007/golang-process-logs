//go:build linux
// +build linux

package processes

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	// . "binalyze-test/types"
)

func getPids() ([]int32, error) {
	pidSlice := make([]int32, 0)

	d, err := os.Open("/proc")
	if err != nil {
		return pidSlice, nil
	}
	defer d.Close()

	for {
		names, err := d.Readdirnames(100)

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		for _, name := range names {

			fmt.Println("initial name", name)
			// Prcoes names start with numbers
			if name[0] < '0' || name[0] > '9' {
				continue
			}

			pid, err := strconv.Atoi(name)

			fmt.Println("name", name, pid)
			// Assumption that process does not exist
			if err != nil {
				continue
			}

			pidSlice = append(pidSlice, int32(pid))
		}
	}

	return pidSlice, nil
}

func statInfo(pid int32) (any, error) {
	fmt.Println("pid", pid)

	statPath := fmt.Sprintf("/proc/%d/stat", pid)

	contents, err := os.ReadFile(statPath)
	if err != nil {
		return nil, nil
	}

	fmt.Println("data", string(contents))

	nameStart := bytes.IndexByte(contents, '(') + 1

	// Filter from start point
	nameEnd := bytes.IndexRune(contents, ')')

	name := string(contents[nameStart+1 : nameEnd])

	fmt.Println("name", name)

	otherStats := strings.Fields(string(contents[nameEnd+2:]))

	fmt.Println("other stats", otherStats)

	return nil, nil
}

func formatLine(line string) []string {
	slice := strings.Fields(line)
	array := []string{}

	i := 0
	for i <= 11 {
		if len(slice) == 0 {
			continue
		} else {
			if len(array) == 10 {

				_string := strings.Join(slice[i:], " ")

				array = append(array, _string)

				break
			} else {
				array = append(array, slice[i])
			}
			i++
		}
	}

	return array
}

func getProcessList() ([]any, error) {
	pids, _ := getPids()

	fmt.Println("v", pids)

	for _, pid := range pids {
		statInfo(pid)
	}

	t := make([]any, 0)

	return t, nil
	// statInfoFromProc()
	// cmd := exec.Command("ps", "aux")
	// out, err := cmd.StdoutPipe()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer out.Close()

	// err = cmd.Start()

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// scanner := bufio.NewScanner(out)

	// processLists := make([][]string, 0)

	// for scanner.Scan() {
	// 	line := scanner.Text()

	// 	processLists = append(processLists, formatLine(line))

	// }

	processes := make([]any, 0)

	// for i, processList := range processLists {

	// 	if i == 0 {
	// 		continue
	// 	}

	// 	process := Process{}

	// 	for i, _process := range processList {
	// 		if i == 0 {
	// 			process.User = _process
	// 		}
	// 		if i == 1 {
	// 			_pid, _ := strconv.Atoi(_process)
	// 			process.PID = int32(_pid)
	// 		}
	// 		if i == 2 {
	// 			cpu_usage, _ := strconv.ParseFloat(_process, 32)
	// 			process.CpuUsage = float32(cpu_usage)
	// 		}
	// 		if i == 3 {
	// 			memory_percentage_usage, _ := strconv.ParseFloat(_process, 32)
	// 			process.MemoryPercentageUsage = float32(memory_percentage_usage)
	// 		}
	// 		if i == 4 {
	// 			virtual_memory_size, _ := strconv.Atoi(_process)
	// 			process.VirtualMemorySize = int64(virtual_memory_size)
	// 		}
	// 		if i == 5 {
	// 			resident_memory_size, _ := strconv.Atoi(_process)
	// 			process.ResidentMemorySize = int64(resident_memory_size)
	// 		}
	// 		if i == 6 {
	// 			process.Tty = _process
	// 		}
	// 		if i == 7 {
	// 			process.State = _process
	// 		}
	// 		if i == 8 {
	// 			process.Started = _process
	// 		}
	// 		if i == 9 {
	// 			process.TotalTime = _process
	// 		}
	// 		if i == 10 {
	// 			process.Command = _process
	// 		}
	// 	}

	// 	processes = append(processes, process)
	// }

	return processes, nil
}
