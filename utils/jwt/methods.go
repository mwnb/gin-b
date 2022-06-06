package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(name, account, extra string) (string, error) {
	now := time.Now()
	expireTime := now.Add(3 * time.Hour)

	jd := JwtData{
		name,
		account,
		extra,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "it is me",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, jd)
	token, err := tokenClaims.SignedString([]byte(secretKey))
	return token, err
}

func ParseToken(token string) (*JwtData, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtData{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JwtData); ok {
			return claims, nil
		}
	}
	return nil, err
}
