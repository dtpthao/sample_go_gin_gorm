package entity

type Contract struct {
	Uuid     string `json:"uuid,omitempty" binding:"required"`
	Name     string `json:"name"`
	UserUuid string `json:"user_uuid" binding:"required"`
	Info     byte   `json:"info"`
}
