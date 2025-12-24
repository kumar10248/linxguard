package config

type Threshold struct {
	Warning  float64 `yaml:"warning"`
	Critical float64 `yaml:"critical"`
}

type Config struct {
	IntervalSeconds int       `yaml:"interval_seconds"`
	CPU             Threshold `yaml:"cpu"`
	Memory          Threshold `yaml:"memory"`
	Disk            Threshold `yaml:"disk"`
	Zombie struct {
		Warning int `yaml:"warning"`
	} `yaml:"zombie"`
}
