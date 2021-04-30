package database

import (
	"log"
	"os"
	"rpiSite/config"
	"rpiSite/models"
	"rpiSite/utils"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	DB   *gorm.DB
	user models.User
)

func Connect() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logLevel(),
			Colorful:      colorful(),
		},
	)

	db, err := gorm.Open(sqlite.Open(config.DB), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		log.Println("Failed to connect database.")
		panic(err)
	}

	DB = db
	log.Println("Database successfully connected.")
}

func Migrate(tables ...interface{}) error {
	log.Println("Migrating database tables.")
	return DB.AutoMigrate(tables...)
}

func Initialize() {
	Connect()

	// Generate data for development.
	if dropTables() && config.IS_DEBUG == "true" {
		log.Println("Dropping database tables.")
		Drop(&user)
		defer Seed()
	}

	Migrate(&user)
}

func Drop(dst ...interface{}) error {
	return DB.Migrator().DropTable(dst...)
}

func Seed() {
	users := []models.User{
		{
			Username: "gusted",
			Email:    "gusted@gusted.xyz",
			Password: utils.GenerateHashedPassword("gusted1234"),
			Role:     models.Admin,
		},
		{
			Username: "john",
			Email:    "john@gusted.xyz",
			Password: utils.GenerateHashedPassword("johnjohn"),
			Role:     models.Regular,
		},
	}

	for _, user := range users {
		DB.Create(&user)
	}
}
