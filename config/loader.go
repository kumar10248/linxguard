package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "/etc/linxguard/linxguard.yml"

func Load() (*Config, error) {
	cfg := defaultConfig()

	data, err := os.ReadFile(configPath)
	if err != nil {
		return cfg, nil // fallback to defaults
	}

	err = yaml.Unmarshal(data, cfg)
	if err != nil {
		return cfg, nil // ignore broken config
	}

	return cfg, nil
}

func defaultConfig() *Config {
	return &Config{
		IntervalSeconds: 15,
		CPU:             Threshold{Warning: 75, Critical: 90},
		Memory:          Threshold{Warning: 75, Critical: 90},
		Disk:            Threshold{Warning: 85, Critical: 95},
		Zombie: struct {
			Warning int `yaml:"warning"`
		}{Warning: 10},
	}
}
