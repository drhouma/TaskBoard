package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Config struct {
	PostgresDSN string `json:"postgres_dsn"`
	UpPort      string `json:"up_port"`
}

func New(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open config file: %w", err)
	}
	defer file.Close()

	body, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("read config file: %w", err)
	}

	var cfg Config
	if err = json.Unmarshal(body, &cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config file: %w", err)
	}

	return &cfg, nil
}
