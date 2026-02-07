package common

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func Go(f func()) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				err := fmt.Errorf("%v", r)
				fmt.Println("panic err: ", err)
				fmt.Sprintf("recovered from panic: %s", string(debug.Stack()))
				return
			}
		}()
	}()
	f()
}

type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int64, secretKey string, expirationTime time.Duration) (string, error) {
	expiration := time.Now().Add(expirationTime)

	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // HS256 is a common signing method

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString, secretKey string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
