package database

import (
	"log"
	"os"
	"rpiSite/config"
	"rpiSite/models"
	"rpiSite/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// DB is the main database variable.
	DB   *gorm.DB
	user models.User
)

// connect to the database.
func connect() {
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
	if !fiber.IsChild() {
		log.Println("Database successfully connected.")
	}
}

// migrate the given tables.
func migrate(tables ...interface{}) error {
	log.Println("Migrating database tables.")
	return DB.AutoMigrate(tables...)
}

// Initialize the database.
func Initialize() {
	connect()

	if !fiber.IsChild() {
		// Generate data for development.
		if dropTables() && config.IsDebug == "true" {
			log.Println("Dropping database tables.")
			if err := drop(&user); err != nil {
				log.Fatalf("Couldn't drop, due to: %v", err.Error())
			}
			defer seed()
		}

		if err := migrate(&user); err != nil {
			log.Fatalf("Couldn't migrate, due to: %v", err.Error())
		}
	}
}

// drop given tables.
func drop(dst ...interface{}) error {
	return DB.Migrator().DropTable(dst...)
}

// seed the database with data.
func seed() {
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
