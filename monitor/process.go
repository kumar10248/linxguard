package monitor

import (
	"os"
	"strconv"
	"strings"
)

// GetZombieCount returns number of zombie processes
func GetZombieCount() (int, error) {
	procEntries, err := os.ReadDir("/proc")
	if err != nil {
		return 0, err
	}

	zombieCount := 0

	for _, entry := range procEntries {
		// process directories are numeric
		if !entry.IsDir() {
			continue
		}

		_, err := strconv.Atoi(entry.Name())
		if err != nil {
			continue
		}

		statPath := "/proc/" + entry.Name() + "/stat"
		data, err := os.ReadFile(statPath)
		if err != nil {
			continue
		}

		fields := strings.Fields(string(data))
		if len(fields) < 3 {
			continue
		}

		// field[2] is process state
		if fields[2] == "Z" {
			zombieCount++
		}
	}

	return zombieCount, nil
}
