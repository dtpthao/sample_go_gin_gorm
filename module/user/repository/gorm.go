package repository

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"sample-go-server/entity"
	"time"
)

const TableUsers = "users"

type User struct {
	Uuid      string                `gorm:"size:36;primaryKey"`
	Username  string                `gorm:"size:64;unique;not null;column:username"`
	Password  string                `gorm:"size:128;column:password"`
	IsAdmin   bool                  //`gorm:"column:is_admin"`
	CreatedAt time.Time             `gorm:"autoCreateTime"`
	UpdatedAt time.Time             `gorm:"autoUpdateTime"`
	IsDeleted soft_delete.DeletedAt `gorm:"size:1;softDelete:flag"`
	db        *gorm.DB
}

func (r User) FromEntity(u entity.User) User {
	return User{
		Uuid:     u.Uuid,
		Username: u.Username,
		Password: u.Password,
		IsAdmin:  u.IsAdmin,
	}
}

func NewUserRepository(db *gorm.DB) entity.IUserRepo {
	return &User{db: db}
}

func (r User) ToEntity() *entity.User {
	return &entity.User{
		Uuid:     r.Uuid,
		Username: r.Username,
		Password: r.Password,
		IsAdmin:  r.IsAdmin,
	}
}

func (r User) Create(u entity.User) (*entity.User, error) {
	tx := r.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	ur := r.FromEntity(u)
	err := tx.Table(TableUsers).Create(&ur).Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	return ur.ToEntity(), err
}

func (r User) GetUserByUsername(username string) (*entity.User, error) {
	var res User
	err := r.db.Table(TableUsers).Where("username = ?", username).Take(&res).Error
	if err != nil {
		return nil, err
	}
	return res.ToEntity(), nil
}

func (r User) GetUserByUuid(uuid string) (*entity.User, error) {
	var res User
	err := r.db.Table(TableUsers).Where("uuid = ?", uuid).Take(&res).Error
	if err != nil {
		return nil, err
	}
	return res.ToEntity(), nil
}

func (r User) Delete(uuid string) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	res := tx.Table(TableUsers).Where("uuid = ?", uuid).Delete(&User{})
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("failed to update user")
	}

	if tx.Commit().Error != nil {
		tx.Rollback()
		return errors.New("cannot commit transaction")
	}

	return nil
}

func (r User) GetList() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Table(TableUsers).Find(&users).Error
	return users, err
}

func (r User) Update(uuid string, data map[string]any) error {

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	res := tx.Table(TableUsers).Where("uuid = ?", uuid).Updates(data)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("failed to update user")
	}

	if tx.Commit().Error != nil {
		tx.Rollback()
		return errors.New("cannot commit transaction")
	}

	return nil
}
