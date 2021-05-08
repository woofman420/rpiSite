package database

import (
	"rpiSite/config"

	"gorm.io/gorm/logger"
)

func logLevel() logger.LogLevel {
	switch config.DBDebug {
	case "error":
		return logger.Error
	case "warn":
		return logger.Warn
	case "info":
		return logger.Info
	default:
		return logger.Silent
	}
}

func colorful() bool {
	switch config.DBColor {
	case "true", "yes", "1":
		return true
	default:
		return false
	}
}

func dropTables() bool {
	switch config.DBDrop {
	case "true", "yes", "1":
		return true
	default:
		return false
	}
}
