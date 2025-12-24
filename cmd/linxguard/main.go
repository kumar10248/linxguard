package main

import (
	"fmt"
	"os"
	"os/exec"

	"linxguard/daemon"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	switch os.Args[1] {
	case "daemon":
		daemon.Start()
	case "status":
		showStatus()
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  linxguard daemon   Start the daemon (used by systemd)")
	fmt.Println("  linxguard status   Show daemon status")
}

func showStatus() {
	cmd := exec.Command("systemctl", "is-active", "linxguard")
	out, err := cmd.Output()

	if err != nil {
		fmt.Println("❌ LinxGuard is not running")
		fmt.Println("👉 Try: sudo systemctl start linxguard")
		return
	}

	if string(out) == "active\n" {
		fmt.Println("✅ LinxGuard is running")
		fmt.Println("📌 Managed by systemd")
	} else {
		fmt.Println("❌ LinxGuard status:", string(out))
	}
}
