package entity

import "time"

type Contract struct {
	Uuid        string    `json:"uuid,omitempty" binding:"required"`
	Name        string    `json:"name"`
	UserUuid    string    `json:"user_uuid" binding:"required"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at,omitempty" time_format:"2006-01-02 15:04:05"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"  time_format:"2006-01-02 15:04:05"`
	IsDeleted   bool      `json:"is_deleted,omitempty"`
}

type NewContractRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type UpdateContractRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (u UpdateContractRequest) ToMap() map[string]any {
	return map[string]any{
		"name":        u.Name,
		"description": u.Description,
	}
}

type UpdateContractResponse struct {
	Success bool `json:"success"`
}

type DeleteContractResponse struct {
	Success bool `json:"success"`
}

//- As a staff, I can create/update/view list/view detail/delete contracts.
//* POST /api/contracts/ (create)
//* GET /api/contracts/ (get list)
//* PATCH/PUT /api/contracts/<id>/ (update)
//* GET /api/contracts/<id>/ (get detail)
//* DELETE /api/contracts/<id>/ (delete)

type IContractUseCase interface {
	CreateNew(uUuid string, c NewContractRequest) (*Contract, error)
	GetListByUser(userUuid string) ([]Contract, error)
	Update(cUuid string, c UpdateContractRequest) error
	GetDetails(cUuid string) (*Contract, error)
	DeleteByUser(cUuid string, uUuid string) error
	Delete(cUuid string) error
}

type IContractRepository interface {
	Add(c Contract) (*Contract, error)
	GetListByUser(userUuid string) ([]Contract, error)
	GetList() ([]Contract, error)
	Update(cUuid string, data any) error
	GetDetails(cUuid string) (*Contract, error)
	DeleteByUser(cUuid string, uUuid string) error
	Delete(cUuid string) error
}
