package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	GRPCServer `yaml:"grpc_server"`
	Postgres   `yaml:"postgres"`
}

type GRPCServer struct {
	Address           string        `yaml:"address"`
	MaxConnectionIdle time.Duration `yaml:"max_connection_idle"`
	Timeout           time.Duration `yaml:"timeout"`
	TimePingClient    time.Duration `yaml:"time_ping_client"`
}

type Postgres struct {
	Connection  string `yaml:"connection"`
	StoragePath string `yaml:"storage_path"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load environment: %w", err)
	}

	appEnv := os.Getenv("APP_ENV")
	if appEnv == "" {
		appEnv = "develop"
	}

	configPath := fmt.Sprintf("configs/%s/config.yaml", appEnv)
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file: %s", err.Error())
	}
	defer file.Close()

	var cfg Config
	if err := yaml.NewDecoder(file).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %s", err.Error())
	}

	return &cfg, nil
}
