package handler

import (
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"glintecoTask/entity"
	usecase3 "glintecoTask/module/contract/usecase"
	handler2 "glintecoTask/module/token/handler"
	usecase2 "glintecoTask/module/token/usecase"
	"glintecoTask/module/user/usecase"
	"glintecoTask/test"
	"net/http"
	"testing"
)

const APIPath = "/api/accounts/"

func SetupRouter() *gin.Engine {
	return gin.Default()
}

func Setup(mockData test.MockData) ContractHandler {
	jwtSecret, _ := hex.DecodeString(test.JWTSecret)
	contractRepo := test.NewTestContractRepo(mockData)
	contractUC := usecase3.NewContractUseCase(contractRepo)

	userRepo := test.NewTestUserRepo(mockData)
	tokenUC := usecase2.NewTokenUseCase(jwtSecret)
	userUC := usecase.NewUserUseCase(userRepo, tokenUC)

	router := SetupRouter()

	tHandler := handler2.NewTokenHandler(tokenUC, userUC, jwtSecret)
	router.Use(tHandler.Authenticate)

	mockKafka := test.MockKafka{}

	contractHandler, _ := NewContractHandler(router, contractUC, mockKafka)

	return contractHandler
}

func TestContractHandler_CreateNew(t *testing.T) {
	mockData := test.NewMockData()
	handler := Setup(mockData)

	tests := []struct {
		name   string
		user   entity.User
		new    *entity.NewContractRequest
		status int
	}{
		{
			"Admin create their own contract",
			mockData.Admin,
			&entity.NewContractRequest{
				Name:        "New Contract",
				UserUuid:    mockData.Admin.Uuid,
				Description: "New contract description",
			},
			http.StatusOK,
		},
		{
			"Admin create contract for staff",
			mockData.Admin,
			&entity.NewContractRequest{
				Name:        "New Contract",
				UserUuid:    mockData.Staff.Uuid,
				Description: "New contract description",
			},
			http.StatusOK,
		},
		{
			"Staff success",
			mockData.Staff,
			&entity.NewContractRequest{
				Name:        "New Contract",
				UserUuid:    mockData.Staff.Uuid,
				Description: "New contract description",
			},
			http.StatusOK,
		},
		{
			"Staff create contract for other",
			mockData.Staff,
			&entity.NewContractRequest{
				Name:        "New Contract",
				UserUuid:    mockData.Admin.Uuid,
				Description: "New contract description",
			},
			http.StatusBadRequest,
		},
		{
			"Invalid request body",
			mockData.Staff,
			nil,
			http.StatusBadRequest,
		},
		{
			"Malicious user",
			mockData.InvalidUser,
			&entity.NewContractRequest{
				Name:        "New Contract",
				Description: "New contract description",
			},
			http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := test.NewMockGinContext()
			mock.SetMiddleware(tt.user)

			mock.Post(tt.new)
			mock.RunTest(handler.CreateNew)

			assert.Equal(t, tt.status, mock.ResponseStatus())

			if tt.status == http.StatusOK {
				var res entity.Contract
				json.Unmarshal(mock.ResponseBody(), &res)
				assert.NotEmpty(t, res)
			}
		})
	}
}

func TestContractHandler_Delete(t *testing.T) {
	mockData := test.NewMockData()
	handler := Setup(mockData)

	tests := []struct {
		name         string
		user         entity.User
		contractUuid string
		status       int
	}{
		{
			"Admin delete admin contract",
			mockData.Admin,
			mockData.AdminContractsUuid()[0],
			http.StatusNoContent,
		},
		{
			"Admin delete staff contract",
			mockData.Admin,
			mockData.StaffContractsUuid()[0],
			http.StatusNoContent,
		},
		{
			"Admin delete non-exist contract",
			mockData.Admin,
			"invalid-contract-uuid",
			http.StatusNoContent,
		},
		{
			"Staff delete their contract",
			mockData.Staff,
			mockData.StaffContractsUuid()[1],
			http.StatusNoContent,
		},
		{
			"Staff delete someone else's contract",
			mockData.Staff,
			mockData.AdminContractsUuid()[1],
			http.StatusNoContent,
		},
		{
			"Staff delete non-exist contract",
			mockData.Staff,
			"invalid-contract-uuid",
			http.StatusNoContent,
		},
		//{ this is middleware test
		//	"Malicious user",
		//	mockData.InvalidUser,
		//	mockData.AdminContractsUuid()[0],
		//	http.StatusBadRequest,
		//},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := test.NewMockGinContext()
			mock.SetMiddleware(tt.user)

			params := []gin.Param{
				{
					Key:   "uuid",
					Value: tt.contractUuid,
				},
			}
			mock.Delete(params)
			mock.RunTest(handler.Delete)

			assert.Equal(t, tt.status, mock.ResponseStatus())
		})
	}
}

func TestContractHandler_GetDetails(t *testing.T) {
	mockData := test.NewMockData()
	handler := Setup(mockData)

	tests := []struct {
		name         string
		user         entity.User
		contractUuid string
		status       int
	}{
		{
			"Admin get admin contract",
			mockData.Admin,
			mockData.AdminContractsUuid()[0],
			http.StatusOK,
		},
		{
			"Admin get staff contract",
			mockData.Admin,
			mockData.StaffContractsUuid()[0],
			http.StatusOK,
		},
		{
			"Staff get their contract",
			mockData.Staff,
			mockData.StaffContractsUuid()[1],
			http.StatusOK,
		},
		{
			"Staff get someone else's contract",
			mockData.Staff,
			mockData.AdminContractsUuid()[1],
			http.StatusBadRequest,
		},
		{
			"Get non-exist contract",
			mockData.Staff,
			"invalid-contract-uuid",
			http.StatusInternalServerError,
		},
		{
			"Malicious user",
			mockData.InvalidUser,
			mockData.AdminContractsUuid()[0],
			http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := test.NewMockGinContext()
			mock.SetMiddleware(tt.user)

			params := []gin.Param{
				{
					Key:   "uuid",
					Value: tt.contractUuid,
				},
			}

			mock.Get(params, nil)
			mock.RunTest(handler.GetDetails)

			assert.Equal(t, tt.status, mock.ResponseStatus())

			if tt.status == http.StatusOK {
				var res entity.Contract
				json.Unmarshal(mock.ResponseBody(), &res)
				assert.NotEmpty(t, res)
			}
		})
	}
}

func TestContractHandler_GetList(t *testing.T) {
	mockData := test.NewMockData()
	handler := Setup(mockData)

	tests := []struct {
		name   string
		user   entity.User
		status int
	}{
		{
			"Admin get admin contracts",
			mockData.Admin,
			http.StatusOK,
		},
		{
			"Staff get their contracts",
			mockData.Staff,
			http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := test.NewMockGinContext()
			mock.SetMiddleware(tt.user)
			mock.Get(nil, nil)
			mock.RunTest(handler.GetList)

			assert.Equal(t, tt.status, mock.ResponseStatus())

			if tt.status == http.StatusOK {
				var res []entity.Contract
				json.Unmarshal(mock.ResponseBody(), &res)
				assert.NotEmpty(t, res)
			}
		})
	}
}

func TestContractHandler_UpdateContract(t *testing.T) {
	mockData := test.NewMockData()
	handler := Setup(mockData)

	tests := []struct {
		name   string
		user   entity.User
		cUuid  string
		req    *entity.UpdateContractRequest
		status int
	}{
		{
			"Admin update admin's contract",
			mockData.Admin,
			mockData.AdminContractsUuid()[0],
			&entity.UpdateContractRequest{
				Name:        "update contract",
				Description: "update description",
			},
			http.StatusOK,
		},
		{
			"Admin update staff's contract",
			mockData.Admin,
			mockData.StaffContractsUuid()[0],
			&entity.UpdateContractRequest{
				Name:        "update contract",
				Description: "update description",
			},
			http.StatusOK,
		},
		{
			"Admin update non-exist contract",
			mockData.Admin,
			"invalid-contract-uuid",
			&entity.UpdateContractRequest{
				Name:        "update contract",
				Description: "update description",
			},
			http.StatusInternalServerError,
		},
		{
			"Staff update their contract",
			mockData.Staff,
			mockData.StaffContractsUuid()[0],
			&entity.UpdateContractRequest{
				Name:        "update contract",
				Description: "update description",
			},
			http.StatusOK,
		},
		{
			"Staff update admin's contract",
			mockData.Staff,
			mockData.AdminContractsUuid()[0],
			&entity.UpdateContractRequest{
				Name:        "update contract",
				Description: "update description",
			},
			http.StatusBadRequest,
		},
		{
			"Staff update non-exist contract",
			mockData.Staff,
			"invalid-uuid",
			&entity.UpdateContractRequest{
				Name:        "update contract",
				Description: "update description",
			},
			http.StatusInternalServerError,
		},
		{
			"Invalid request body",
			mockData.Staff,
			mockData.StaffContractsUuid()[0],
			nil,
			http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock := test.NewMockGinContext()
			mock.SetMiddleware(tt.user)

			params := []gin.Param{
				{
					Key:   "uuid",
					Value: tt.cUuid,
				},
			}

			mock.Put(tt.req, params)
			mock.RunTest(handler.UpdateContract)

			assert.Equal(t, tt.status, mock.ResponseStatus())
		})
	}
}
