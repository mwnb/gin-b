package jwt

import "github.com/dgrijalva/jwt-go"

type JwtData struct {
	Name    string `json:"name"`
	Account string `json:"account"`
	Extra   string `json:"extra"`
	jwt.StandardClaims
}
