package monitor

import (
	"syscall"
)

// DiskUsage represents disk usage stats
type DiskUsage struct {
	Total uint64
	Used  uint64
	Free  uint64
	Usage float64
}

// GetDiskUsage returns disk usage for a given path (e.g. "/")
func GetDiskUsage(path string) (DiskUsage, error) {
	var stat syscall.Statfs_t

	err := syscall.Statfs(path, &stat)
	if err != nil {
		return DiskUsage{}, err
	}

	total := stat.Blocks * uint64(stat.Bsize)
	free := stat.Bavail * uint64(stat.Bsize)
	used := total - free

	usage := (float64(used) / float64(total)) * 100

	return DiskUsage{
		Total: total,
		Used:  used,
		Free:  free,
		Usage: usage,
	}, nil
}
