package delivery

import (
	"github.com/gin-gonic/gin"
	"glintecoTask/entity"
)

type ContractHandler struct {
	uc     entity.IContractUseCase
	router *gin.RouterGroup
}

func NewContractHandler(app *gin.Engine, uc entity.IContractUseCase) ContractHandler {
	return ContractHandler{
		uc:     uc,
		router: app.Group("/api/contracts"),
	}
}

func (h ContractHandler) RegisterHandler(middlewares ...gin.HandlerFunc) {
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
