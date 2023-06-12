package adapter

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type EncryptAdapter struct{}

func NewEncryptAdapter() *EncryptAdapter {
	return &EncryptAdapter{}
}

func (e EncryptAdapter) GenerateToken(payload map[string]interface{}, minutesToExpire uint) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(minutesToExpire)).Unix()
	for k, v := range payload {
		claims[k] = v
	}

	var secretKey = os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	return tokenString
}
