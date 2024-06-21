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

func (uc UserUseCase) GetList() ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uc UserUseCase) Update(u entity.User) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (uc UserUseCase) Delete(username string) error {
	return uc.repo.Delete(username)
}

func NewUserUseCase(r entity.IUserRepo, tuc entity.ITokenUseCase) entity.IUserUseCase {
	return UserUseCase{
		repo:    r,
		tokenUC: tuc,
	}
}

func (uc UserUseCase) Create(u entity.User) (*entity.User, error) {
	_, err := uc.repo.GetDetails(u.Username)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		// todo hash password
		return uc.repo.Create(u)
	}
	return nil, err
}

func (uc UserUseCase) Login(u entity.User) (string, error) {

	dbUser, err := uc.repo.GetDetails(u.Username)
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

func (uc UserUseCase) GetDetails(username string) (*entity.User, error) {
	return uc.repo.GetDetails(username)
}

func (uc UserUseCase) Logout(username string) error {
	// TODO implement me
	panic("implement me")
}
