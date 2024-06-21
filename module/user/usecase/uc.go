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
	return uc.repo.GetList()
}

func (uc UserUseCase) Update(u entity.User) error {
	data := map[string]any{ // When update with struct, GORM will only update non-zero fields, you might want to use map to update attributes or use Select to specify fields to update
		"username": u.Username,
		"password": u.Password,
		"is_admin": u.IsAdmin,
	}
	return uc.repo.Update(u.Uuid, data)
}

func (uc UserUseCase) Delete(userUuid string) error {
	return uc.repo.Delete(userUuid)
}

func NewUserUseCase(r entity.IUserRepo, tuc entity.ITokenUseCase) entity.IUserUseCase {
	return UserUseCase{
		repo:    r,
		tokenUC: tuc,
	}
}

func (uc UserUseCase) Create(u entity.User) (*entity.User, error) {
	_, err := uc.repo.GetUserByUsername(u.Username)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return uc.repo.Create(u)
	}

	if err == nil {
		return nil, errors.New("user already exist")
	}
	return nil, err
}

func (uc UserUseCase) Login(u entity.User) (string, error) {

	dbUser, err := uc.repo.GetUserByUsername(u.Username)
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

func (uc UserUseCase) GetUserByUsername(username string) (*entity.User, error) {
	return uc.repo.GetUserByUsername(username)
}

func (uc UserUseCase) GetUserByUuid(uuid string) (*entity.User, error) {
	return uc.repo.GetUserByUuid(uuid)
}

func (uc UserUseCase) Logout(username string) error {
	// TODO implement me
	panic("implement me")
}
