package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"sync"
	"time"
)

var (
	jwtSecret []byte
	once      sync.Once
)

func getSecret() ([]byte, error) {
	var err error
	once.Do(func() {
		secret := os.Getenv("SECRET")
		if secret == "" {
			err = errors.New("JWT секрет не задан в переменной окружения SECRET")
			return
		}
		jwtSecret = []byte(secret)
	})
	return jwtSecret, err
}

func GenerateJWT(userID int64) (string, error) {
	secret, err := getSecret()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	return token.SignedString(secret)
}
