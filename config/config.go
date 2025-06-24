package config

import (
	"fmt"
	"os"
)

type Config struct {
	TemporalNamespace string
	TemporalAddress   string
	VaultAddress      string
	VaultToken        string
	GeminiAPIKey      string
	TaskQueue         string
}

func Load() (*Config, error) {
	config := &Config{
		TemporalNamespace: getEnvOrDefault("TEMPORAL_NAMESPACE", "default"),
		TemporalAddress:   getEnvOrDefault("TEMPORAL_ADDRESS", "localhost:7233"),
		VaultAddress:      getEnvOrDefault("VAULT_ADDR", "http://localhost:8200"),
		VaultToken:        os.Getenv("VAULT_TOKEN"),
		GeminiAPIKey:      os.Getenv("GEMINI_API_KEY"),
		TaskQueue:         getEnvOrDefault("TASK_QUEUE", "ADFLOWHQ_TASK_QUEUE"),
	}

	// Validate required environment variables
	if config.VaultToken == "" {
		return nil, fmt.Errorf("VAULT_TOKEN environment variable is required")
	}

	return config, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
