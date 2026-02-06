package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerPort     string
	LogLevel       string
	KubeConfigPath string
	CacheTTL       int
	EnableMockData bool
}

func Load() *Config {
	return &Config{
		ServerPort:     getEnv("SERVER_PORT", "8080"),
		LogLevel:       getEnv("LOG_LEVEL", "info"),
		KubeConfigPath: getEnv("KUBECONFIG", ""),
		CacheTTL:       getEnvAsInt("CACHE_TTL_SECONDS", 300),
		EnableMockData: getEnvAsBool("ENABLE_MOCK_DATA", true),
	}
}

func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return fallback
}

func getEnvAsBool(key string, fallback bool) bool {
	if value := os.Getenv(key); value != "" {
		if boolVal, err := strconv.ParseBool(value); err == nil {
			return boolVal
		}
	}
	return fallback
}
