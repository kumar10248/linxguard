package monitor

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// MemoryUsage holds RAM stats
type MemoryUsage struct {
	Total uint64
	Used  uint64
	Free  uint64
	Usage float64
}

// GetMemoryUsage reads /proc/meminfo and calculates usage
func GetMemoryUsage() (MemoryUsage, error) {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		return MemoryUsage{}, err
	}
	defer file.Close()

	var memTotal, memAvailable uint64

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "MemTotal:") {
			fields := strings.Fields(line)
			memTotal, _ = strconv.ParseUint(fields[1], 10, 64)
		}

		if strings.HasPrefix(line, "MemAvailable:") {
			fields := strings.Fields(line)
			memAvailable, _ = strconv.ParseUint(fields[1], 10, 64)
		}
	}

	// values are in KB
	used := memTotal - memAvailable
	usage := (float64(used) / float64(memTotal)) * 100

	return MemoryUsage{
		Total: memTotal * 1024,
		Used:  used * 1024,
		Free:  memAvailable * 1024,
		Usage: usage,
	}, nil
}
