package authenticator

import "github.com/golang-jwt/jwt"

type MyClaims struct {
	jwt.StandardClaims
	UserId   string
	Username string `json:"Username"`
	Email    string `json:"Email"`
}
