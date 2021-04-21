package config

import "os"

var (
	IS_DEBUG = getEnv("DEBUG", "false")
	PORT     = getEnv("PORT", ":8080")
)

func getEnv(value, backupValue string) string {
	if envValue, ok := os.LookupEnv(value); ok {
		return envValue
	}
	return backupValue
}
