🛡️ LinxGuard

LinxGuard is a beginner-friendly Linux system guardian daemon written in Go.
It continuously monitors system health and explains problems in simple human language, instead of throwing cryptic metrics at users.

## 📦 Installation

Download the latest `.deb` from GitHub Releases:
https://github.comkumar10248/linxguard/releases


Built with a systems-engineering mindset: safe by default, non-root, systemd-managed, and production-ready.

✨ Features

🧠 Human-readable explanations (not just numbers)

⚙️ systemd-managed daemon

🔐 Runs as a non-root system user

🚨 Alert deduplication (no alert spam)

📊 Real-time monitoring:

CPU usage

Memory usage

Disk usage

Zombie processes

🧩 Clean architecture (monitor → rules → explain)

📦 Single static Go binary

🖥️ What LinxGuard Monitors
Component	Method
CPU	/proc/stat
Memory	/proc/meminfo
Disk	syscall.Statfs
Processes	/proc/[pid]/stat
Service lifecycle	systemd

No shell parsing. No unsafe commands.

📸 Example Output
⚠️ High Memory Usage [WARNING]

🧠 What’s happening:
Available RAM is running low and the system may slow down.

👉 Suggested action:
Close unused applications or check memory-hungry processes.


Alerts are shown only when state changes, not repeatedly.

🧱 Architecture
CLI (linxguard)
   ↓
Daemon (systemd)
   ↓
Monitors (CPU / Memory / Disk / Zombies)
   ↓
Rules Engine
   ↓
Explanation Engine


This separation keeps the system maintainable and testable.

🚀 Installation
1️⃣ Build
go build -o linxguard ./cmd/linxguard
sudo mv linxguard /usr/local/bin/

2️⃣ Create system user
sudo useradd --system --no-create-home --shell /usr/sbin/nologin linxguard

3️⃣ Install systemd service
sudo cp systemd/linxguard.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable linxguard
sudo systemctl start linxguard

4️⃣ Verify
linxguard status
systemctl status linxguard

🧪 Usage
Start daemon (used by systemd)
linxguard daemon

Check status
linxguard status

View logs
journalctl -u linxguard -f

🔐 Security Design

Runs as a dedicated system user

No root privileges at runtime

systemd hardening:

NoNewPrivileges

ProtectSystem

ProtectHome

PrivateTmp

Read-only system inspection

LinxGuard never modifies your system by default.

🧠 Why LinxGuard?

Most monitoring tools:

Assume expert users

Show raw metrics

Spam alerts

Require complex setup

LinxGuard focuses on:

Beginners

Clarity

Safety

Zero confusion

🎯 Who Is This For?

Linux beginners

Students learning system internals

Developers running Linux locally

Anyone who wants “what’s wrong?” answered clearly

📌 Interview Talking Points

You can confidently say:

“I built a Linux daemon in Go managed by systemd.”

“It monitors CPU, memory, disk, and zombie processes using /proc.”

“I implemented a rules engine to convert metrics into human explanations.”

“The daemon runs as a non-root user with systemd hardening.”

“I added alert deduplication to prevent monitoring noise.”

🔥 These are strong system-engineering signals.

🛣️ Roadmap

 Config file support (/etc/linxguard.yml)

 .deb package

 Optional safe auto-fix mode

 Network monitoring

 TUI dashboard

📄 License

MIT License

🤝 Contributing

PRs welcome.
Ideas, bug reports, and feature requests are encouraged.


