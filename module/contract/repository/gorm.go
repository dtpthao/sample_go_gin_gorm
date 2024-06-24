package repository

import (
	"errors"
	"glintecoTask/entity"
	"glintecoTask/module/user/repository"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
	"time"
)

const TableContract = "contracts"

type Contract struct {
	Uuid        string                `gorm:"size:36;primaryKey;column:uuid"`
	Name        string                `gorm:"size:255;column:name"`
	UserUuid    string                `gorm:"size:36;not null;index"`
	Description string                `gorm:"column:description"`
	CreatedAt   time.Time             `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt   time.Time             `gorm:"autoCreateTime;column:updated_at"`
	IsDeleted   soft_delete.DeletedAt `gorm:"size:1;softDelete:flag"`
	User        repository.User       `gorm:"foreignKey:UserUuid;references:Uuid;OnUpdate:CASCADE;OnDelete:CASCADE"`
	db          *gorm.DB
}

func NewContractRepository(db *gorm.DB) entity.IContractRepository {
	return Contract{db: db}
}

func (r Contract) FromEntity(c entity.Contract) Contract {
	return Contract{
		Uuid:        c.Uuid,
		Name:        c.Name,
		UserUuid:    c.UserUuid,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
		UpdatedAt:   c.UpdatedAt,
	}
}

func (r Contract) ToEntity() *entity.Contract {
	return &entity.Contract{
		Uuid:        r.Uuid,
		Name:        r.Name,
		UserUuid:    r.UserUuid,
		Description: r.Description,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
	}
}

func (r Contract) Add(c entity.Contract) (*entity.Contract, error) {
	cr := r.FromEntity(c)
	err := r.db.Table(TableContract).Create(&cr).Error
	return cr.ToEntity(), err
}

func (r Contract) GetListByUser(uUuid string) ([]entity.Contract, error) {
	var contracts []entity.Contract
	err := r.db.Table(TableContract).Where("user_uuid = ?", uUuid).Find(&contracts).Error
	return contracts, err
}

func (r Contract) GetList() ([]entity.Contract, error) {
	var contracts []entity.Contract
	err := r.db.Table(TableContract).Find(&contracts).Error
	return contracts, err
}

func (r Contract) Update(cUuid string, data any) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	res := tx.Table(TableContract).Where("uuid = ?", cUuid).Updates(data)
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

func (r Contract) GetDetails(cUuid string) (*entity.Contract, error) {
	var res Contract
	err := r.db.Table(TableContract).Where("uuid = ?", cUuid).Take(&res).Error
	if err != nil {
		return nil, err
	}
	return res.ToEntity(), nil
}

func (r Contract) Delete(cUuid string) error {
	c := Contract{Uuid: cUuid}

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := r.db.Table(TableContract).Delete(&c).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	if tx.Commit().Error != nil {
		tx.Rollback()
		return errors.New("cannot commit transaction")
	}

	return nil
}

func (r Contract) DeleteByUser(cUuid string, uUuid string) error {
	c := Contract{Uuid: cUuid}

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	res := r.db.Table(TableContract).Where("uuid = ? and user_uuid = ?", cUuid, uUuid).Delete(&c)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	if res.RowsAffected == 0 {
		tx.Rollback()
		return errors.New("delete contract failed")
	}

	if tx.Commit().Error != nil {
		tx.Rollback()
		return errors.New("cannot commit transaction")
	}

	return nil
}
