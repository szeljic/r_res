package common

import "github.com/dgrijalva/jwt-go"

var JwtKey []byte

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var smtpHost string
var smtpPort string