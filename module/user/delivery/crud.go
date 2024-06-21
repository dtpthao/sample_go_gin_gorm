package delivery

import (
	"github.com/gin-gonic/gin"
	"glintecoTask/entity"
	"net/http"
)

type UserLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (ul UserLogin) ToEntity() entity.User {
	return entity.User{
		Username: ul.Username,
		Password: ul.Password,
	}
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h UserHandler) Login(c *gin.Context) {
	var userRegister UserLogin
	err := c.ShouldBindJSON(&userRegister)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	// todo need to invalidate token when logout
	token, err := h.uc.Login(userRegister.ToEntity())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, LoginResponse{Token: token})
}

func (h UserHandler) CreateUser(c *gin.Context) {
	var userReq entity.User
	err := c.ShouldBindJSON(&userReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	user, err := h.uc.Create(userReq)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h UserHandler) GetListUsers(c *gin.Context) {}

func (h UserHandler) UpdateUser(c *gin.Context) {}

func (h UserHandler) GetUserDetail(c *gin.Context) {}

func (h UserHandler) DeleteUser(c *gin.Context) {}

func (h UserHandler) Logout(c *gin.Context) {
	// todo
}
