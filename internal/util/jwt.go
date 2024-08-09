package util

import (
	"errors"
	"time"

	"github.com/evertonbzr/api-golang/internal/config"
	"github.com/evertonbzr/api-golang/internal/model"
	"github.com/golang-jwt/jwt/v5"
)

type ModuleClaims struct {
	jwt.RegisteredClaims
	User model.User `json:"user"`
}

func HasJwtExpired(token *jwt.Token) error {
	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		return err
	}

	if exp == nil || exp.Before(time.Now()) {
		return errors.New("TokenHasExpired")
	}

	return nil
}

func GetDurationFromJWT(token *jwt.Token) (time.Duration, error) {
	exp, err := token.Claims.GetExpirationTime()
	if err != nil {
		return 0, err
	}
	return time.Since(exp.Time).Abs(), nil
}

func DecodeJWT(tokenString string) (*jwt.Token, *ModuleClaims, error) {
	claims := &ModuleClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("UnexpectedSigningMethod")
		}

		return []byte(config.JWT_SECRET), nil
	})
	if err != nil {
		return nil, nil, err
	}

	if err = HasJwtExpired(token); err != nil {
		return nil, nil, err
	}

	return token, claims, nil
}

func GenerateJwt(user *model.User) (string, error) {
	expTime := time.Now().Add(8760 * time.Hour)

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		ModuleClaims{
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expTime),
			},
			User: *user,
		})

	tokenString, err := token.SignedString([]byte(config.JWT_SECRET))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
