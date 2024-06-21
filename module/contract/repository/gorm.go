package repository

import (
	"errors"
	"glintecoTask/entity"
	"gorm.io/gorm"
	"time"
)

const TableContract = "contract"

type ContractRepository struct {
	Uuid      string    `gorm:"primaryKey;column:uuid"`
	Name      string    `gorm:"column:name"`
	UserUuid  string    `gorm:"column:user_uuid"`
	Details   any       `gorm:"column:details"`
	CreatedAt time.Time `gorm:"autoCreateTime;column:created_at"`
	UpdatedAt time.Time `gorm:"autoCreateTime;column:updated_at"`
	db        *gorm.DB
}

func NewContractRepository(db *gorm.DB) entity.IContractRepository {
	return ContractRepository{db: db}
}

func (r ContractRepository) FromEntity(c entity.Contract) ContractRepository {
	return ContractRepository{
		Uuid:      c.Uuid,
		Name:      c.Name,
		UserUuid:  c.UserUuid,
		Details:   c.Details,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func (r ContractRepository) ToEntity() *entity.Contract {
	return &entity.Contract{
		Uuid:      r.Uuid,
		Name:      r.Name,
		UserUuid:  r.UserUuid,
		Details:   r.Details,
		CreatedAt: r.CreatedAt,
		UpdatedAt: r.UpdatedAt,
	}
}

func (r ContractRepository) Add(c entity.Contract) (*entity.Contract, error) {
	cr := r.FromEntity(c)
	err := r.db.Table(TableContract).Create(&cr).Error
	return cr.ToEntity(), err
}

func (r ContractRepository) GetListByUser(uUuid string) ([]entity.Contract, error) {
	var contracts []entity.Contract
	err := r.db.Table(TableContract).Where("user_uuid = ?", uUuid).Find(&contracts).Error
	return contracts, err
}

func (r ContractRepository) GetList() ([]entity.Contract, error) {
	var contracts []entity.Contract
	err := r.db.Table(TableContract).Find(&contracts).Error
	return contracts, err
}

func (r ContractRepository) Update(cUuid string, data any) error {
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

func (r ContractRepository) GetDetails(cUuid string) (*entity.Contract, error) {
	var res ContractRepository
	err := r.db.Table(TableContract).Where("uuid = ?", cUuid).Take(&res).Error
	if err != nil {
		return nil, err
	}
	return res.ToEntity(), nil
}

func (r ContractRepository) Delete(cUuid string) error {
	c := ContractRepository{Uuid: cUuid}

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

func (r ContractRepository) DeleteByUser(cUuid string, uUuid string) error {
	c := ContractRepository{Uuid: cUuid}

	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	err := r.db.Table(TableContract).Where("uuid = ? and user_uuid = ?", cUuid, uUuid).Delete(&c).Error
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
