package usecase

import "glintecoTask/entity"

type ContractUseCase struct{}

func (ContractUseCase) CreateNew(c entity.NewContractRequest) (*entity.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func (ContractUseCase) GetList() ([]entity.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func (ContractUseCase) Update(c entity.UpdateContractRequest) error {
	//TODO implement me
	panic("implement me")
}

func (ContractUseCase) GetInfo(cUuid string) (*entity.Contract, error) {
	//TODO implement me
	panic("implement me")
}

func (ContractUseCase) Delete(cUuid string) error {
	//TODO implement me
	panic("implement me")
}
