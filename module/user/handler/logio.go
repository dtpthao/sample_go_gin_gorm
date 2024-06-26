package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sample-go-server/entity"
	"sample-go-server/utils"
)

func (h UserHandler) Login(c *gin.Context) {
	var userRegister entity.UserLogin
	err := c.ShouldBindJSON(&userRegister)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}

	// todo need to invalidate token when logout
	token, err := h.uc.Login(userRegister.ToEntity())
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, entity.LoginResponse{Token: token})
}
