package models

import (
	"errors"
	"rpiSite/config"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Role int

const (
	Regular Role = iota
	Moderator
	Admin
)

type User struct {
	gorm.Model `json:"-"`
	Username   string `gorm:"unique;not null" validate:"required,username,min=5,max=20"`
	Email      string `gorm:"unique" validate:"required,email"`
	Password   string `validate:"required,min=8,max=32"`
	Role       Role   `gorm:"default=0"`
}

type APIUser struct {
	Username string
	Email    string
	ID       uint
	Role     Role
}

func getDBSession(db *gorm.DB) (tx *gorm.DB) {
	if config.DB_DEBUG == "info" {
		return db.Session(&gorm.Session{
			Logger: db.Logger.LogMode(logger.Info),
		})
	} else {
		return db.Session(&gorm.Session{
			Logger: db.Logger.LogMode(logger.Silent),
		})
	}
}

func FindUserByEmail(db *gorm.DB, email string) (*User, error) {
	user := new(User)

	if res := db.Where("email = ?", email).First(&user); res.Error != nil {
		return nil, res.Error
	}

	if user.ID == 0 {
		return nil, errors.New("user not found")
	}

	return user, nil
}

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
		return nil, errors.New("user not found")
	}

	return user, nil
}
