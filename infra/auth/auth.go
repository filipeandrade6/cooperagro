package auth

import (
	"github.com/golang-jwt/jwt"
)

// jwtCustomClaims are custom claims extending default ones.
// See https://github.com/golang-jwt/jwt for more examples
type JWTCustomClaims struct {
	ID    string   `json:"name"`
	Roles []string `json:"roles"`
	jwt.StandardClaims
}
