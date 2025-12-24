package daemon

import (
	"log"
	"time"

	"linxguard/config"
	"linxguard/explain"
	"linxguard/logger"
	"linxguard/monitor"
	"linxguard/rules"
)

// activeAlerts remembers emitted alerts to avoid spamming
var activeAlerts = make(map[string]bool)

// emitOnce prints an alert only when it appears or changes
func emitOnce(key string, issue *rules.Issue) {
	if issue == nil {
		delete(activeAlerts, key) // issue resolved
		return
	}

	if activeAlerts[key] {
		return // already reported
	}

	explain.Print(issue)
	activeAlerts[key] = true
}

func Start() {
	// Initialize logging first
	logger.Init()

	// Load configuration (safe fallback to defaults)
	cfg, _ := config.Load()
	log.Printf("Config loaded: interval=%ds", cfg.IntervalSeconds)

	// Handle SIGTERM / SIGINT
	HandleSignals()

	log.Println("linxguard daemon started")

	for {
		// ---------- CPU ----------
		cpu, err := monitor.GetCPUUsage()
		if err != nil {
			log.Println("CPU monitor error:", err)
		} else {
			log.Printf("CPU Usage: %.2f%%", cpu)
			emitOnce(
				"cpu",
				rules.CheckCPUWithThreshold(cpu, cfg.CPU.Warning, cfg.CPU.Critical),
			)
		}

		// ---------- Disk ----------
		disk, err := monitor.GetDiskUsage("/")
		if err != nil {
			log.Println("Disk monitor error:", err)
		} else {
			log.Printf("Disk Usage: %.2f%%", disk.Usage)
			emitOnce("disk", rules.CheckDisk(disk.Usage))
		}

		// ---------- Memory ----------
		mem, err := monitor.GetMemoryUsage()
		if err != nil {
			log.Println("Memory monitor error:", err)
		} else {
			log.Printf("Memory Usage: %.2f%%", mem.Usage)
			emitOnce("memory", rules.CheckMemory(mem.Usage))
		}

		// ---------- Zombies ----------
		zombies, err := monitor.GetZombieCount()
		if err != nil {
			log.Println("Zombie monitor error:", err)
		} else {
			log.Printf("Zombie Processes: %d", zombies)
			emitOnce("zombie", rules.CheckZombies(zombies))
		}

		// Sleep based on config
		time.Sleep(time.Duration(cfg.IntervalSeconds) * time.Second)
	}
}
