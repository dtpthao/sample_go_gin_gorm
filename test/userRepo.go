package test

import (
	"errors"
	"gorm.io/gorm"
	"sample-go-server/entity"
)

type UserRepository struct {
	Data MockData
}

func NewTestUserRepo(m MockData) entity.IUserRepo {
	return &UserRepository{Data: m}
}

func (ur UserRepository) Create(u entity.User) (*entity.User, error) {
	return &u, nil
}

func (ur UserRepository) GetUserByUsername(username string) (*entity.User, error) {

	switch username {
	case ur.Data.Admin.Username:
		return &ur.Data.Admin, nil
	case ur.Data.Staff.Username:
		return &ur.Data.Staff, nil
	default:
		return nil, gorm.ErrRecordNotFound
	}
}

func (ur UserRepository) GetUserByUuid(uuid string) (*entity.User, error) {
	switch uuid {
	case ur.Data.Admin.Uuid:
		return &ur.Data.Admin, nil
	case ur.Data.Staff.Uuid:
		return &ur.Data.Staff, nil
	default:
		return nil, gorm.ErrRecordNotFound
	}
}

func (ur UserRepository) Delete(uuid string) error {
	switch uuid {
	case ur.Data.Admin.Uuid, ur.Data.Staff.Uuid:
		return nil
	default:
		return errors.New("delete user failed")
	}
}

func (ur UserRepository) GetList() ([]entity.User, error) {
	return []entity.User{ur.Data.Admin, ur.Data.Staff}, nil
}

func (ur UserRepository) Update(uuid string, data map[string]any) error {
	switch uuid {
	case ur.Data.Admin.Uuid, ur.Data.Staff.Uuid:
		return nil
	default:
		return errors.New("update user failed")
	}
}
