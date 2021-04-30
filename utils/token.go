package utils

import (
	"rpiSite/config"
	"time"

	"github.com/form3tech-oss/jwt-go"
)

type JWTTokenBuilder struct {
	*jwt.Token
}

func NewJWTToken() *JWTTokenBuilder {
	return &JWTTokenBuilder{jwt.New(jwt.SigningMethodHS512)}
}

func (jt *JWTTokenBuilder) SetClaim(name string, value interface{}) *JWTTokenBuilder {
	jt.Claims.(jwt.MapClaims)[name] = value
	return jt
}

func (jt *JWTTokenBuilder) SetExpiration(duration time.Time) *JWTTokenBuilder {
	if !duration.IsZero() {
		jt.Claims.(jwt.MapClaims)["exp"] = duration.Unix()
	}
	return jt
}

func (jt *JWTTokenBuilder) GetSignedString(customKey []byte) (string, error) {
	if customKey == nil {
		customKey = []byte(config.JWT_SIGNING_KEY)
	}
	return jt.SignedString(customKey)
}
