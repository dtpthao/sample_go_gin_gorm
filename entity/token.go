package entity

import (
	"time"
)

type Token struct {
	Username string    `json:"user_uuid,omitempty"`
	Exp      time.Time `json:"exp,omitempty"`
	Iat      time.Time `json:"iat,omitempty"`
}

type ITokenUseCase interface {
	Create(u User) (string, error)
	Verify(token string) (*Token, error)
}
