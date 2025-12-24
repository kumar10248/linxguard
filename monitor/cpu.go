package monitor

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"time"
)

// CPUStat holds CPU time values
type CPUStat struct {
	idle  uint64
	total uint64
}

// readCPUStat reads CPU stats from /proc/stat
func readCPUStat() (CPUStat, error) {
	file, err := os.Open("/proc/stat")
	if err != nil {
		return CPUStat{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()

	fields := strings.Fields(scanner.Text())

	var idle, total uint64
	for i, v := range fields {
		if i == 0 {
			continue // skip "cpu"
		}
		val, _ := strconv.ParseUint(v, 10, 64)
		total += val
		if i == 4 { // idle field
			idle = val
		}
	}

	return CPUStat{idle: idle, total: total}, nil
}

// GetCPUUsage returns CPU usage percentage
func GetCPUUsage() (float64, error) {
	stat1, err := readCPUStat()
	if err != nil {
		return 0, err
	}

	time.Sleep(1 * time.Second)

	stat2, err := readCPUStat()
	if err != nil {
		return 0, err
	}

	idleDelta := stat2.idle - stat1.idle
	totalDelta := stat2.total - stat1.total

	usage := 100 * (1 - float64(idleDelta)/float64(totalDelta))
	return usage, nil
}
