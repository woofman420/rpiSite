package utils

import (
	"time"

	"rpiSite/config"

	"github.com/form3tech-oss/jwt-go"
)

// JWTTokenBuilder is wrapper around *jwt.Token but more chain-able.
type JWTTokenBuilder struct {
	*jwt.Token
}

// NewJWTToken returns an empty HS512 JWT.
func NewJWTToken() *JWTTokenBuilder {
	return &JWTTokenBuilder{jwt.New(jwt.SigningMethodHS512)}
}

// SetClaim will set a key,value.
func (jt *JWTTokenBuilder) SetClaim(name string, value interface{}) *JWTTokenBuilder {
	jt.Claims.(jwt.MapClaims)[name] = value
	return jt
}

// SetExpiration the duration the token is valid.
func (jt *JWTTokenBuilder) SetExpiration(duration time.Time) *JWTTokenBuilder {
	if !duration.IsZero() {
		jt.Claims.(jwt.MapClaims)["exp"] = duration.Unix()
	}
	return jt
}

// GetSignedString will returned the signed string.
func (jt *JWTTokenBuilder) GetSignedString(customKey []byte) (string, error) {
	if customKey == nil {
		customKey = UnsafeByteConversion(config.JWTSigningKey)
	}
	return jt.SignedString(customKey)
}
