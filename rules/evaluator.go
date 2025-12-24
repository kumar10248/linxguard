package rules

func CheckCPU(cpu float64) *Issue {
	if cpu > 90 {
		return &Issue{
			Severity: Critical,
			Title:    "High CPU Usage",
			Explanation: "Your system CPU is heavily loaded. Applications may become unresponsive.",
			Suggestion:  "Close unused applications or check for runaway processes using top or htop.",
		}
	}
	if cpu > 75 {
		return &Issue{
			Severity: Warning,
			Title:    "Elevated CPU Usage",
			Explanation: "CPU usage is higher than normal and may affect performance.",
			Suggestion:  "Monitor running applications and consider closing heavy tasks.",
		}
	}
	return nil
}

func CheckDisk(diskUsage float64) *Issue {
	if diskUsage > 95 {
		return &Issue{
			Severity: Critical,
			Title:    "Disk Almost Full",
			Explanation: "Disk space is nearly exhausted. This can break updates and system services.",
			Suggestion:  "Delete unused files or clear logs in /var/log.",
		}
	}
	if diskUsage > 85 {
		return &Issue{
			Severity: Warning,
			Title:    "High Disk Usage",
			Explanation: "Disk usage is high and may cause system instability.",
			Suggestion:  "Clean up large files or move data to external storage.",
		}
	}
	return nil
}

func CheckZombies(count int) *Issue {
	if count > 10 {
		return &Issue{
			Severity: Warning,
			Title:    "Zombie Process Accumulation",
			Explanation: "Zombie processes indicate that parent processes are not cleaning up properly.",
			Suggestion:  "Restart the parent application or reboot if the issue persists.",
		}
	}
	return nil
}

func CheckMemory(memUsage float64) *Issue {
	if memUsage > 90 {
		return &Issue{
			Severity: Critical,
			Title:    "Memory Almost Exhausted",
			Explanation: "Available RAM is critically low. The system may freeze or start killing processes.",
			Suggestion:  "Close heavy applications or consider adding more RAM or swap.",
		}
	}
	if memUsage > 75 {
		return &Issue{
			Severity: Warning,
			Title:    "High Memory Usage",
			Explanation: "Memory usage is high and may slow down the system.",
			Suggestion:  "Close unused applications or check memory-hungry processes.",
		}
	}
	return nil
}

func CheckCPUWithThreshold(cpu, warn, crit float64) *Issue {
	if cpu > crit {
		return &Issue{
			Severity: Critical,
			Title:    "High CPU Usage",
			Explanation: "CPU usage is critically high and the system may become unresponsive.",
			Suggestion:  "Close heavy applications or investigate runaway processes.",
		}
	}
	if cpu > warn {
		return &Issue{
			Severity: Warning,
			Title:    "Elevated CPU Usage",
			Explanation: "CPU usage is higher than normal and may affect performance.",
			Suggestion:  "Monitor running applications and consider closing heavy tasks.",
		}
	}
	return nil
}
