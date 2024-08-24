// go:build linux
// go:build linux
//go:build linux
// +build linux

package process

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	"binalyze-test/utils"

	. "binalyze-test/types"
)

type Memory struct {
	MemTotal     int
	MemFree      int
	MemAvailable int
}

// Fetch clock time as it will be used for calulations

var clkTick float64 = 100

func init() {
	// Running init function
	if stdout, err := exec.Command("getconf", "CLK_TCK").Output(); err == nil {

		stdoutArray := strings.Fields(string(stdout))
		if val, err := strconv.ParseFloat(stdoutArray[0], 64); err == nil {
			clkTick = val
		}
	}
}

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

			// Prcoes names start with numbers
			if name[0] < '0' || name[0] > '9' {
				continue
			}

			pid, err := strconv.Atoi(name)
			// Assumption that process does not exist
			if err != nil {
				continue
			}

			pidSlice = append(pidSlice, int32(pid))
		}
	}

	log.Println("Linux pids fetched successfully")

	return pidSlice, nil
}

func getUserName(pid int32) (string, error) {
	statPath := fmt.Sprintf("/proc/%d/status", pid)

	contents, err := os.ReadFile(statPath)
	if err != nil {
		return "", err
	}
	status := strings.Split(string(contents), "\n")

	var uid string

	for _, stat := range status {
		tabParts := strings.SplitN(stat, "\t", -1)

		if len(tabParts) < 2 {
			continue
		}
		value := tabParts[1]

		switch strings.TrimRight(tabParts[0], ":") {
		case "Uid":
			uids := strings.Split(value, "\t")

			uid = uids[0]
		}
	}

	user, err := user.LookupId(uid)
	if err != nil {
		return "", err
	}

	return user.Username, nil
}

func getMemorySize() (int64, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return 0, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var n int
		if nItems, _ := fmt.Sscanf(scanner.Text(), "MemTotal: %d kB", &n); nItems == 1 {
			return int64(n), nil
		}
	}

	return 1, nil
}

/*
This function takes in uTime , sTime and pStartTime
uTime: time the program used in the user mode(in clock time)
sTime: time the program used in the kernel mode(in clock time)
cuTime: time the program used in the user's code
totalCpuTime  = uTime + sTime + cuTime
pStartTime: time since the process started(in clock time)
Fetch the uptime of the system(Time since the system booted)
Fetch the clock tick for the OS (defaults 100)
convert the uTime , sTime and pStartTime to seconds

eTime : Time elapsed since programmed started and pc booted
totalCpuUsageSec: Total time used by the process in both user mode and kernel mode
cpuUsage = (ratio of totalCpuUsageSec to eTime) * 100

*/

func calcCpuUsage(totalCpuTime, pStartTime float64) (float64, error) {
	var cpuUsage float64

	// Get system uptime

	stdout, err := exec.Command("cat", "/proc/uptime").Output()
	if err != nil {
		return cpuUsage, err
	}

	uptime := strings.Fields(string(stdout))

	sysUptime, err := strconv.ParseFloat(uptime[0], 64)
	if err != nil {
		return cpuUsage, err
	}

	totalCpuTimeSec := totalCpuTime / clkTick

	pStartTimeSec := pStartTime / clkTick

	eTime := sysUptime - pStartTimeSec

	cpuUsage = (totalCpuTimeSec * 100) / eTime

	return cpuUsage, nil
}

/*
This function takes in rss and return memeory percentage
rss : Size of process in the memory
Get the total physical memory with meminfo

memoryUsage = (ratio of rss to total phycial memeory) * 100

NOTE: This function might no the actual memory, as we need to calculate the memory across all processor
This calculation comes from her https://man7.org/linux/man-pages/man1/ps.1.html

*/

func calcMemoryUsage(rss float64) (float64, error) {
	var memoryUsage float64

	cmd := exec.Command("cat", "/proc/meminfo")

	output, err := cmd.StdoutPipe()
	if err != nil {
		return memoryUsage, err
	}

	defer output.Close()

	memTotal, err := getMemorySize()
	if err != nil {
		return memoryUsage, err
	}

	memoryUsage = rss * 100.00 / float64(memTotal)

	return memoryUsage, nil
}

func statInfo(p *Process, pid int32) error {
	statPath := fmt.Sprintf("/proc/%d/stat", pid)

	contents, err := os.ReadFile(statPath)
	if err != nil {
		return nil
	}

	nameStart := bytes.IndexByte(contents, '(') + 1

	// Filter from start point
	nameEnd := bytes.IndexRune(contents, ')')

	name := string(contents[nameStart:nameEnd])

	stats := strings.Fields(string(contents[nameEnd+2:]))

	p.PID = pid

	username, err := getUserName(pid)
	if err != nil {
		return err
	}

	p.User = username

	// time a programm used in user mode in seconds
	uTime, err := strconv.ParseFloat(stats[11], 64)
	if err != nil {
		return err
	}

	// time a programm used in kernel mode in seconds
	sTime, err := strconv.ParseFloat(stats[12], 64)
	if err != nil {
		return err
	}

	cuTime, err := strconv.ParseFloat(stats[13], 64)
	if err != nil {
		return err
	}

	// Process start time since system booted
	pStartTime, err := strconv.ParseFloat(stats[19], 64)
	if err != nil {
		return err
	}
	totalCpuTime := uTime + sTime + cuTime
	cpuUsage, err := calcCpuUsage(totalCpuTime, pStartTime)
	if err != nil {
		return err
	}

	p.CpuUsage = utils.FormatTo2Decimal(cpuUsage)

	residentMemorySize, err := strconv.ParseUint(stats[21], 10, 64)
	if err != nil {
		return err
	}

	p.ResidentMemorySize = int64(residentMemorySize / 1000)

	memoryUsage, err := calcMemoryUsage(float64(residentMemorySize))
	if err != nil {
		return err
	}

	p.MemoryUsage = utils.FormatTo2Decimal(memoryUsage)

	virtualMemorySize, err := strconv.ParseUint(stats[20], 10, 64)
	if err != nil {
		return err
	}
	p.VirtualMemorySize = int64(virtualMemorySize / 1000)

	p.State = formatState(stats[0])

	p.TotalTime = fmt.Sprintf("%.2f", pStartTime/clkTick)

	p.CpuTime = fmt.Sprintf("%.2f", (uTime+sTime)/clkTick)

	p.Command = name

	priority, err := strconv.Atoi(stats[16])
	if err != nil {
		return err
	}

	p.Priority = guagePriority(priority)

	return nil
}
