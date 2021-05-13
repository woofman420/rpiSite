package utils

import (
	"errors"
	"fmt"
)

var (
	// ErrorUserNotFound returns an user has not been found error.
	ErrorUserNotFound           = errors.New("user not found")
	errorUnexpectedJWTAlgoirthm = errors.New("unexpected jwt signing method")
	errorCantMigrate            = errors.New("couldn't migrate, due to")
	errorCantDrop               = errors.New("couldn't drop, due to")
)

// ErrorUnexpectedJWTAlgoirthm returns that a unexpected Signing method was found.
func ErrorUnexpectedJWTAlgoirthm(errorMessage string) error {
	return fmt.Errorf("%w: %s", errorUnexpectedJWTAlgoirthm, errorMessage)
}

// ErrorCantMigrate returns that it couldn't migate.
func ErrorCantMigrate(errorMessage string) error {
	return fmt.Errorf("%w: %s", errorCantMigrate, errorMessage)
}

// ErrorCantDrop returns that it couldn't drop.
func ErrorCantDrop(errorMessage string) error {
	return fmt.Errorf("%w: %s", errorCantDrop, errorMessage)
}
