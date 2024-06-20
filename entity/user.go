package entity

import (
	"time"
)

type User struct {
	Uuid      string    `json:"uuid,omitempty"`
	Username  string    `json:"username" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	IsAdmin   bool      `json:"is_admin,default=false"`
	Active    bool      `json:"active,default=true"`
	CreatedAt time.Time `json:"created_at,omitempty" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `json:"updated_at,omitempty"  time_format:"2006-01-02 15:04:05"`
}

type IUserRepo interface {
	Create(u User) error
	FindByUsername(username string) (*User, error)
	DeleteByUsername(username string) error
}

type IUserUseCase interface {
	Register(u User) error
	Login(u User) (string, error)
	Logout(username string) error
}
