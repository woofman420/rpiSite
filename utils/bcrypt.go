package utils

import (
	"log"
	"rpiSite/config"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

var salt = getSalt()

func getSalt() int {
	salt, err := strconv.Atoi(config.SALT)
	if err != nil {
		log.Fatalln("Failed to convert SALT env variable, err:", err)
	}

	return salt
}

func GenerateHashedPassword(pw string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pw), salt)
	if err != nil {
		log.Println(err)
	}

	return string(hash)
}

func CompareHashedPassword(user, form string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user), []byte(form))
	if err != nil {
		return err
	}

	return nil
}
