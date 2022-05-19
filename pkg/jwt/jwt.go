package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	// Secret key being used to sign tokens
	SecretKey = []byte("secret")
)

// GenerateToken generates a JWT token, assigns a username to it's claims and retuns it.
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Fatalf("failed type casting")
	}
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatalf("failed signing the jwt with secret: %v", err)
		return "", err
	}
	return tokenString, nil
}

// ParseToken parses a JWT token and returns the username in it's claims.
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
