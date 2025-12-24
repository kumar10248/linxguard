package explain

import (
	"fmt"

	"linxguard/rules"
)

func Print(issue *rules.Issue) {
	if issue == nil {
		return
	}

	fmt.Println("────────────────────────────")
	fmt.Printf("⚠️  %s [%s]\n", issue.Title, issue.Severity)
	fmt.Println()
	fmt.Println("🧠 What’s happening:")
	fmt.Println(issue.Explanation)
	fmt.Println()
	fmt.Println("👉 Suggested action:")
	fmt.Println(issue.Suggestion)
	fmt.Println("────────────────────────────")
}
