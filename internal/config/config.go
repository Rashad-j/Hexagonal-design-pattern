package config

import (
	"os"
)

type Config struct {
	Port    string
	GinMode string
}

func LoadConfig() *Config {
	return &Config{
		Port:    getEnv("PORT", "8080"),
		GinMode: getEnv("GIN_MODE", "debug"),
	}
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
