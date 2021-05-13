package models

import (
	"rpiSite/config"
	"rpiSite/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Role is an alias for int and is equivalent to int in all ways. It is
// used to represent a specifc role.
type Role int

const (
	// Regular role.
	Regular Role = iota
	// Moderator role.
	Moderator
	// Admin role.
	Admin
)

// User struct that's stored in the database.
type User struct {
	gorm.Model `json:"-"`
	Username   string `gorm:"unique;not null" validate:"required,username,min=5,max=20"`
	Email      string `gorm:"unique" validate:"required,email"`
	Password   string `validate:"required,min=8,max=32"`
	Role       Role   `gorm:"default=0"`
}

// APIUser without the database model.
type APIUser struct {
	Username string
	Email    string
	ID       uint
	Role     Role
}

func getDBSession(db *gorm.DB) (tx *gorm.DB) {
	if config.DBDebug == "info" {
		return db.Session(&gorm.Session{
			Logger: db.Logger.LogMode(logger.Info),
		})
	}
	return db.Session(&gorm.Session{
		Logger: db.Logger.LogMode(logger.Silent),
	})
}

// FindUserByEmail will find the user based of the email within the given database.
func FindUserByEmail(db *gorm.DB, email string) (*User, error) {
	user := new(User)

	if res := db.Where("email = ?", email).First(&user); res.Error != nil {
		return nil, res.Error
	}

	if user.ID == 0 {
		return nil, utils.ErrorUserNotFound
	}

	return user, nil
}

// FindUserByName will find the user based of the username within the given database.
func FindUserByName(db *gorm.DB, name string) (*User, error) {
	user := new(User)

	err := getDBSession(db).
		Where("username = ?", name).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, utils.ErrorUserNotFound
	}

	return user, nil
}
