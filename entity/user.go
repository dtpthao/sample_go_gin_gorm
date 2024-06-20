package entity

import (
	"time"
)

type User struct {
	Uuid        string    `json:"uuid,omitempty"`
	Username    string    `json:"username" binding:"required"`
	DisplayName string    `json:"display_name,omitempty"`
	Email       string    `json:"email" binding:"required,email"`
	Role        int       `json:"role,default=0" binding:"gte=0,lte=1"`
	Active      bool      `json:"active_status,default=true"`
	CreatedAt   time.Time `json:"created_at,omitempty" time_format:"2006-01-02 15:04:05"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"  time_format:"2006-01-02 15:04:05"`
}

type IUserRepo interface {
	Create(u User) (User, error)
	Update(u User) (User, error)
	FindByUsername(username string) (User, error)
	DeleteByUsername(username string) error
}
