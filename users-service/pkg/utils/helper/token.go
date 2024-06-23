package helper

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

type TokenUserCase interface {
	GenerateAccessToken(claims JwtCustomClaims) (string, error)
}

type JwtCustomClaims struct {
	UserID   string `json:"user_id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

type TokenUserCaseImpl struct{}

func NewTokenUseCase() *TokenUserCaseImpl {
	return &TokenUserCaseImpl{}
}

// func untuk generator token (siapa yang lagi masuk, siapa yang lagi request)
func (t *TokenUserCaseImpl) GenerateAccessToken(claims JwtCustomClaims) (string, error) {
	plainToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	encodedToken, err := plainToken.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", err
	}

	return encodedToken, nil
}
