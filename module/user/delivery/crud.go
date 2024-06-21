package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"glintecoTask/entity"
	"glintecoTask/utils"
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
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}

	// todo need to invalidate token when logout
	token, err := h.uc.Login(userRegister.ToEntity())
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, LoginResponse{Token: token})
}

type NewUserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	IsAdmin  bool   `json:"is_admin"`
}

func (h UserHandler) CreateUser(c *gin.Context) {
	var userReq NewUserRequest
	err := c.ShouldBindJSON(&userReq)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}

	// todo validate input
	newUser := entity.User{
		Uuid:     uuid.New().String(),
		Username: userReq.Username,
		Password: userReq.Password, // todo hash password from frontend?
		IsAdmin:  userReq.IsAdmin,
		Active:   true,
	}

	user, err := h.uc.Create(newUser)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h UserHandler) GetListUsers(c *gin.Context) {
	//res, err := h.uc.GetList()

}

func (h UserHandler) UpdateUser(c *gin.Context) {}

func (h UserHandler) GetUserDetail(c *gin.Context) {}

func (h UserHandler) DeleteUser(c *gin.Context) {}

func (h UserHandler) Logout(c *gin.Context) {
	// todo
}
