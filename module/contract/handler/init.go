package handler

import (
	"github.com/gin-gonic/gin"
	"glintecoTask/entity"
)

const DeleteContractTopic = "delete-contract"

type ContractHandler struct {
	uc     entity.IContractUseCase
	router *gin.RouterGroup
	kafka  entity.IKafkaService
}

func NewContractHandler(app *gin.Engine, uc entity.IContractUseCase, kafkaService entity.IKafkaService) (ContractHandler, error) {

	kafkaConsumerHandler := NewKafkaContractHandler(uc)
	err := kafkaService.Init([]string{DeleteContractTopic}, kafkaConsumerHandler)
	if err != nil {
		return ContractHandler{}, err
	}

	go kafkaService.Listen()

	return ContractHandler{
		uc:     uc,
		router: app.Group("/api/contracts"),
		kafka:  kafkaService,
	}, nil
}

func (h ContractHandler) RegisterHandler(middlewares ...gin.HandlerFunc) {

	h.router.Use(middlewares...)
	//- As a staff, I can create/update/view list/view detail/delete contracts.
	//* POST /api/contracts/ (create)
	//* GET /api/contracts/ (get list)
	//* PATCH/PUT /api/contracts/<id>/ (update)
	//* GET /api/contracts/<id>/ (get detail)
	//* DELETE /api/contracts/<id>/ (delete)
	h.router.POST("/", h.CreateNew)
	h.router.GET("/", h.GetList)
	h.router.PUT("/:uuid", h.UpdateContract)
	h.router.GET("/:uuid", h.GetDetails)
	h.router.DELETE("/:uuid", h.Delete)
}
