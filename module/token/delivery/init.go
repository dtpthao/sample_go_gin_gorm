package delivery

import "glintecoTask/entity"

type TokenHandler struct {
	uc        *entity.ITokenUseCase
	JWTSecret []byte
}

func NewTokenHandler(uc *entity.ITokenUseCase, jwtSecret []byte) TokenHandler {
	return TokenHandler{
		uc:        uc,
		JWTSecret: jwtSecret,
	}
}
