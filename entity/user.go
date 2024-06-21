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
	Create(u User) (*User, error)
	GetUserByUsername(username string) (*User, error)
	GetUserByUuid(uuid string) (*User, error)
	Delete(uuid string) error
	GetList() ([]User, error)
	Update(uuid string, data map[string]any) error
}

type IUserUseCase interface {
	Create(u User) (*User, error)
	Login(u User) (string, error)
	Logout(uuid string) error
	GetUserByUsername(username string) (*User, error)
	GetUserByUuid(uuid string) (*User, error)
	GetList() ([]User, error)
	Update(u User) error
	Delete(uuid string) error
}
