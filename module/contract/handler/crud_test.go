package handler

import (
	"bytes"
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
	"net/http/httptest"
	"testing"
)

const APIPath = "/api/accounts/"

func SetupRouter() *gin.Engine {
	return gin.Default()
}

func Setup(mockData test.MockData) (*gin.Engine, entity.ITokenUseCase, ContractHandler) {
	jwtSecret, _ := hex.DecodeString(test.JWTSecret)
	contractRepo := test.NewTestContractRepo(mockData)
	contractUC := usecase3.NewContractUseCase(contractRepo)

	userRepo := test.NewTestUserRepo(mockData)
	tokenUC := usecase2.NewTokenUseCase(jwtSecret)
	userUC := usecase.NewUserUseCase(userRepo, tokenUC)

	router := SetupRouter()

	thandler := handler2.NewTokenHandler(tokenUC, userUC, jwtSecret)
	router.Use(thandler.Authenticate)

	return router, tokenUC, NewContractHandler(router, contractUC)
}

func TestContractHandler_CreateNew(t *testing.T) {
	mockData := test.NewMockData()
	router, tokenUC, handler := Setup(mockData)

	router.POST(APIPath, handler.CreateNew)

	tests := []struct {
		name   string
		user   entity.User
		new    *entity.NewContractRequest
		status int
	}{
		{
			"Admin Success",
			mockData.Admin,
			&entity.NewContractRequest{
				Name:        "New Contract",
				Description: "New contract description",
			},
			http.StatusOK,
		},
		{
			"Staff success",
			mockData.Staff,
			&entity.NewContractRequest{
				Name:        "New Contract",
				Description: "New contract description",
			},
			http.StatusOK,
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
			token, _ := tokenUC.Create(tt.user)
			token = "Bearer " + token

			var jsonValue []byte
			if tt.new != nil {
				jsonValue, _ = json.Marshal(tt.new)
			}
			req, _ := http.NewRequest("POST", APIPath, bytes.NewBuffer(jsonValue))
			req.Header.Set("Authorization", token)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)

			if tt.status == http.StatusOK {
				var res entity.Contract
				json.Unmarshal(w.Body.Bytes(), &res)
				assert.NotEmpty(t, res)
			}
		})
	}
}

func TestContractHandler_Delete(t *testing.T) {
	mockData := test.NewMockData()
	router, tokenUC, handler := Setup(mockData)

	router.DELETE(APIPath+":uuid", handler.Delete)

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
			http.StatusInternalServerError,
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
			http.StatusInternalServerError,
		},
		{
			"Staff delete non-exist contract",
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
			token, _ := tokenUC.Create(tt.user)
			token = "Bearer " + token

			req, _ := http.NewRequest("DELETE", APIPath+tt.contractUuid, nil)
			req.Header.Set("Authorization", token)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}
}

func TestContractHandler_GetDetails(t *testing.T) {
	mockData := test.NewMockData()
	router, tokenUC, handler := Setup(mockData)

	router.GET(APIPath+":uuid", handler.GetDetails)

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
			token, _ := tokenUC.Create(tt.user)
			token = "Bearer " + token

			req, _ := http.NewRequest("GET", APIPath+tt.contractUuid, nil)
			req.Header.Set("Authorization", token)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)

			if tt.status == http.StatusOK {
				var res entity.Contract
				json.Unmarshal(w.Body.Bytes(), &res)
				assert.NotEmpty(t, res)
			}
		})
	}
}

func TestContractHandler_GetList(t *testing.T) {
	mockData := test.NewMockData()
	router, tokenUC, handler := Setup(mockData)

	router.GET(APIPath, handler.GetList)

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
		{
			"Malicious user",
			mockData.InvalidUser,
			http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, _ := tokenUC.Create(tt.user)
			token = "Bearer " + token

			req, _ := http.NewRequest("GET", APIPath, nil)
			req.Header.Set("Authorization", token)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)

			if tt.status == http.StatusOK {
				var res []entity.Contract
				json.Unmarshal(w.Body.Bytes(), &res)
				assert.NotEmpty(t, res)
			}
		})
	}
}

func TestContractHandler_UpdateContract(t *testing.T) {
	mockData := test.NewMockData()
	router, tokenUC, handler := Setup(mockData)

	router.PUT(APIPath+":uuid", handler.UpdateContract)

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
			"Staff update their's contract",
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
		{
			"Malicious user",
			mockData.InvalidUser,
			mockData.StaffContractsUuid()[0],
			&entity.UpdateContractRequest{
				Name:        "update contract",
				Description: "update description",
			},
			http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, _ := tokenUC.Create(tt.user)
			token = "Bearer " + token

			var jsonValue []byte
			if tt.req != nil {
				jsonValue, _ = json.Marshal(tt.req)
			}
			req, _ := http.NewRequest("PUT", APIPath+tt.cUuid, bytes.NewBuffer(jsonValue))
			req.Header.Set("Authorization", token)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.status, w.Code)
		})
	}
}
