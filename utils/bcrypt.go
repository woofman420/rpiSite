package utils

import (
	"log"
	"rpiSite/config"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

var salt = getSalt()

func getSalt() int {
	salt, err := strconv.Atoi(config.Salt)
	if err != nil {
		log.Fatalln("Failed to convert SALT env variable, err:", err)
	}

	return salt
}

// GenerateHashedPassword Generate a bcrypt salted password.
func GenerateHashedPassword(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword(UnsafeByteConversion(pwd), salt)
	if err != nil {
		return ""
	}

	return UnsafeStringConversion(hash)
}

// CompareHashedPassword will compare if the given password matches with the hashed password.
func CompareHashedPassword(user, form string) bool {
	return bcrypt.CompareHashAndPassword(UnsafeByteConversion(user), UnsafeByteConversion(form)) != nil
}
