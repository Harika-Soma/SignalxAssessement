package jwt

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	SecretKey = []byte("AIzaSyBtLc5L8WECfgz6i1NzzNU7uFfhIig7DyQ")
)

func GenerateToken(sessionID string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sID"] = sessionID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in Generating key", err)
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	fmt.Println("token parsed successfully", token)
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		sessionID := claims["sID"].(string)
		return sessionID, nil
	} else {
		return "", err
	}
}
