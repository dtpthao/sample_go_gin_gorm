package delivery

import (
	"github.com/gin-gonic/gin"
	"glintecoTask/entity"
)

type UserHandler struct {
	uc     entity.IUserUseCase
	router *gin.RouterGroup
}

func NewUserHandler(app *gin.Engine, uc entity.IUserUseCase) UserHandler {
	return UserHandler{
		uc:     uc,
		router: app.Group("/accounts/"),
	}
}

func (h UserHandler) RegisterHandler() {
	h.router.POST("/login", h.Login)
	//h.router.GET("/logout", h.Logout)
}

func (h UserHandler) RegisterMiddleware(middleware func(c *gin.Context)) {
	h.router.Use(middleware)
}
