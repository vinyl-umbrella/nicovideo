package config

import "os"

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DBName   string
}

var Config = DbConfig{
	User:     getEnv("DB_USER", "user"),
	Password: getEnv("DB_PASSWORD", "password"),
	Host:     getEnv("DB_HOST", "localhost"),
	Port:     getEnv("DB_PORT", "3306"),
	DBName:   getEnv("DB_NAME", "nicovideodb"),
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
