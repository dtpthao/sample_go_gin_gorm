package usecase

import (
	"errors"
	"glintecoTask/entity"
	"gorm.io/gorm"
)

type UserUseCase struct {
	repo    entity.IUserRepo
	tokenUC entity.ITokenUseCase
}

func (uc UserUseCase) Register(u entity.User) error {
	_, err := uc.repo.FindByUsername(u.Username)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// todo hash password
		err = uc.repo.Create(u)
		return err
	}
	return err
}

func (uc UserUseCase) Login(u entity.User) (string, error) {

	dbUser, err := uc.repo.FindByUsername(u.Username)
	if err != nil {
		return "", err
	}

	if dbUser.Password != u.Password {
		return "", errors.New("wrong username or password")
	}

	token, err := uc.tokenUC.Create(*dbUser)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uc UserUseCase) Logout(username string) error {
	// TODO implement me
	panic("implement me")
}

func NewUserUseCase(r entity.IUserRepo, tuc entity.ITokenUseCase) entity.IUserUseCase {
	return UserUseCase{
		repo:    r,
		tokenUC: tuc,
	}
}
