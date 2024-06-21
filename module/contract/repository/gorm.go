package repository

import "glintecoTask/entity"

type ContractRepository struct{}

func (ContractRepository) Add(c entity.Contract) (*entity.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func (ContractRepository) GetList() ([]entity.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func (ContractRepository) Update(c entity.Contract) error {
	//TODO implement me
	panic("implement me")
}

func (ContractRepository) GetInfo(cUuid string) (*entity.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func (ContractRepository) Delete(cUuid string) error {
	//TODO implement me
	panic("implement me")
}
