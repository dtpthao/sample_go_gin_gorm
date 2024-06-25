package handler

import (
	"github.com/gin-gonic/gin"
	"sample-go-server/entity"
)

type UserHandler struct {
	uc     entity.IUserUseCase
	router *gin.RouterGroup
}

func NewUserHandler(app *gin.Engine, uc entity.IUserUseCase) UserHandler {
	return UserHandler{
		uc:     uc,
		router: app.Group("/"),
	}
}

func (h UserHandler) RegisterHandler(middlewares ...gin.HandlerFunc) {
	h.router.POST("accounts/login", h.Login)

	//- As a admin, I can create/update/view list/view detail/delete contracts and staffs.
	//* POST /api/staffs/ (create)
	//* GET /api/staffs/ (get list)
	//* PATCH/PUT /api/staffs/<id>/ (update)
	//* GET /api/staffs/<id>/ (get detail)
	//* DELETE /api/staffs/<id>/ (delete)
	admin := h.router.Group("api/staffs")
	admin.Use(middlewares...)
	admin.POST("/", h.CreateUser)
	admin.GET("/", h.GetListUsers)
	admin.PUT("/:uuid", h.UpdateUser)
	admin.GET("/:uuid", h.GetUserInfo)
	admin.DELETE("/:uuid", h.DeleteUser)
}
