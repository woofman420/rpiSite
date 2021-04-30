package config

import "os"

var (
	IS_DEBUG = getEnv("DEBUG", "false")
	PORT     = getEnv("PORT", ":8080")

	DB       = getEnv("DB", "dev.db")
	DB_DEBUG = getEnv("DB_DEBUG", "false")
	DB_COLOR = getEnv("DB_COLOR", "false")
	DB_DROP  = getEnv("DB_DROP", "false")

	JWT_SIGNING_KEY = getEnv("JWT_SIGNING_KEY", "ABigSecretPassword")
	SALT            = getEnv("SALT", "10")
)

func getEnv(value, backupValue string) string {
	if envValue, ok := os.LookupEnv(value); ok {
		return envValue
	}
	return backupValue
}
