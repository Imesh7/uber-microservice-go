package api

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	UserID int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func CreateToken(user User, expiryTime time.Duration) (string, error) {
	jwtPassword := os.Getenv("JWT_KEY")
	var jwtKey = []byte(jwtPassword)
	claims := &CustomClaims{
		UserID: 1,
		Email:  "jj",
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(expiryTime).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token,&CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		/* _, ok := token.Method.(*jwt.SigningMethodECDSA)
		if !ok {
			return nil, jwt.ValidationError{Errors: 401}
		} */
		return []byte(os.Getenv("JWT_KEY")), nil
	})

}
