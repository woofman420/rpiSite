package config

import "os"

var (
	// IsDebug checks if current runtime is debug env.
	IsDebug = getEnv("DEBUG", "false") == "true"

	// IsProduction checks if current runtime is producntion env.
	IsProduction = !IsDebug

	// Port which the server is listening.
	Port = getEnv("PORT", ":8080")

	// DB file name.
	DB = getEnv("DB", "dev.db")
	// DBDebug to check if Debug mode should be enabled.
	DBDebug = getEnv("DB_DEBUG", "false")
	// DBColor to check if it should color the output.
	DBColor = getEnv("DB_COLOR", "false")
	// DBDrop to seed more data.
	DBDrop = getEnv("DB_DROP", "false")

	// JWTSigningKey is the key that JWT is signing with.
	JWTSigningKey = getEnv("JWT_SIGNING_KEY", "ABigSecretPassword")
	// Salt length.
	Salt = getEnv("SALT", "10")
	// SecretCode to register a new account.
	SecretCode = getEnv("SECRET_CODE", "scawry")
)

func getEnv(value, backupValue string) string {
	if envValue, ok := os.LookupEnv(value); ok {
		return envValue
	}
	return backupValue
}
