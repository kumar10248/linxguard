package rules

type Severity string

const (
	Info     Severity = "INFO"
	Warning  Severity = "WARNING"
	Critical Severity = "CRITICAL"
)

type Issue struct {
	Severity    Severity
	Title       string
	Explanation string
	Suggestion  string
}
