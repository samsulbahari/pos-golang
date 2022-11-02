package libraries

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var signatureKey = []byte("authpos")

func GenerateJWT(userId int, role int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"role":    role,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(signatureKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecryptJwt(token string) (map[string]interface{}, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("auth invalid")
		}
		return signatureKey, nil
	})

	if err != nil || !parsedToken.Valid {
		return map[string]interface{}{}, errors.New("auth invalid")
	}
	return parsedToken.Claims.(jwt.MapClaims), nil
}
