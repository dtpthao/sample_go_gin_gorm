package usecase

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"glintecoTask/entity"
	"time"
)

type TokenUseCase struct {
	jwtSecret []byte
}

func NewTokenUseCase(jwtSecret []byte) entity.ITokenUseCase {
	return &TokenUseCase{
		jwtSecret: jwtSecret,
	}
}

func (uc TokenUseCase) Create(u entity.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Username,
		"exp":      time.Now().Add(time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	})

	tokenString, err := claims.SignedString(uc.jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (uc TokenUseCase) Verify(tokenString string) (*entity.Token, error) {
	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return uc.jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("cannot parse token")
	}

	username, ok := claims["username"].(string)
	if !ok {
		return nil, errors.New("invalid token")
	}

	return &entity.Token{Username: username}, nil
}
