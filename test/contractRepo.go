package test

import (
	"errors"
	"glintecoTask/entity"
	"slices"
)

type ContractRepository struct {
	Data MockData
}

func NewTestContractRepo(m MockData) entity.IContractRepository {
	var cr ContractRepository
	cr.Data = m
	return &ContractRepository{Data: m}
}

func (cr ContractRepository) Add(c entity.Contract) (*entity.Contract, error) {
	switch c.UserUuid {
	case AdminUuid, StaffUuid:
		return &c, nil
	default:
		return nil, errors.New("test failed: user not found")
	}
}

func (cr ContractRepository) GetListByUser(userUuid string) ([]entity.Contract, error) {
	switch userUuid {
	case AdminUuid:
		return cr.Data.AdminContracts, nil
	case StaffUuid:
		return cr.Data.StaffContracts, nil
	default:
		return nil, errors.New("test failed: user not found")
	}
}

func (cr ContractRepository) GetList() ([]entity.Contract, error) {
	var res []entity.Contract
	res = append(res, cr.Data.AdminContracts...)
	res = append(res, cr.Data.StaffContracts...)
	return res, nil
}

func (cr ContractRepository) Update(cUuid string, data any) error {
	switch {
	case slices.Contains(cr.Data.AdminContractsUuid(), cUuid), slices.Contains(cr.Data.StaffContractsUuid(), cUuid):
		return nil
	default:
		return errors.New("contract uuid not found")
	}
}

func (cr ContractRepository) GetDetails(cUuid string) (*entity.Contract, error) {
	idx := slices.Index(cr.Data.AdminContractsUuid(), cUuid)
	if idx >= 0 {
		return &cr.Data.AdminContracts[idx], nil
	}

	idx = slices.Index(cr.Data.StaffContractsUuid(), cUuid)
	if idx >= 0 {
		return &cr.Data.StaffContracts[idx], nil
	}

	return nil, errors.New("contract not found")
}

func (cr ContractRepository) DeleteByUser(cUuid string, uUuid string) error {

	switch uUuid {
	case cr.Data.Admin.Uuid:
		idx := slices.Index(cr.Data.AdminContractsUuid(), cUuid)
		if idx < 0 {
			return errors.New("record not found")
		}
		return nil
	case cr.Data.Staff.Uuid:
		idx := slices.Index(cr.Data.StaffContractsUuid(), cUuid)
		if idx < 0 {
			return errors.New("record not found")
		}
		return nil
	default:
		return errors.New("test failed: user not found")
	}
}

func (cr ContractRepository) Delete(cUuid string) error {
	if slices.Contains(cr.Data.AdminContractsUuid(), cUuid) || slices.Contains(cr.Data.StaffContractsUuid(), cUuid) {
		return nil
	}
	return errors.New("contract not found")
}
