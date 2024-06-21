package entity

type Contract struct {
	Uuid     string `json:"uuid,omitempty" binding:"required"`
	Name     string `json:"name"`
	UserUuid string `json:"user_uuid" binding:"required"`
	Details  any    `json:"details"`
}

type NewContractRequest struct {
	Name     string `json:"name" binding:"required"`
	UserUuid string `json:"user_uuid" binding:"required"`
	Details  any    `json:"details"`
}

type UpdateContractRequest struct {
	Name    string `json:"name"`
	Details any    `json:"details"`
}

//- As a staff, I can create/update/view list/view detail/delete contracts.
//* POST /api/contracts/ (create)
//* GET /api/contracts/ (get list)
//* PATCH/PUT /api/contracts/<id>/ (update)
//* GET /api/contracts/<id>/ (get detail)
//* DELETE /api/contracts/<id>/ (delete)

type IContractUseCase interface {
	CreateNew(c NewContractRequest) (*Contract, error)
	GetList() ([]Contract, error)
	Update(c UpdateContractRequest) error
	GetInfo(cUuid string) (*Contract, error)
	Delete(cUuid string) error
}

type IContractRepository interface {
	Add(c Contract) (*Contract, error)
	GetList() ([]Contract, error)
	Update(c Contract) error
	GetInfo(cUuid string) (*Contract, error)
	Delete(cUuid string) error
}
