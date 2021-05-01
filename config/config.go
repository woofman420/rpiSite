package config

import "os"

var (
	IS_DEBUG = getEnv("DEBUG", "false")
	PORT     = getEnv("PORT", ":8080")

	ADMIN_USER = getEnv("ADMIN_USER", "admin")
	ADMIN_PWD  = getEnv("ADMIN_PWD", "admin")

	JWT_SIGNING_KEY = getEnv("JWT_SIGNING_KEY", "ABigSecretPassword")
	SECRET_CODE     = getEnv("SECRET_CODE", "scawry")
)

func getEnv(value, backupValue string) string {
	if envValue, ok := os.LookupEnv(value); ok {
		return envValue
	}
	return backupValue
}
