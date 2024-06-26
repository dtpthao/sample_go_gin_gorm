package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sample-go-server/entity"
	"sample-go-server/utils"
	"strings"
)

type TokenHandler struct {
	tuc       entity.ITokenUseCase
	uuc       entity.IUserUseCase
	JWTSecret []byte
}

func NewTokenHandler(tuc entity.ITokenUseCase, uuc entity.IUserUseCase, jwtSecret []byte) TokenHandler {
	return TokenHandler{
		tuc:       tuc,
		uuc:       uuc,
		JWTSecret: jwtSecret,
	}
}

func (h TokenHandler) Authenticate(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

	token, err := h.tuc.Verify(tokenString)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}

	user, err := h.uuc.GetUserByUsername(token.Username)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}

	c.Set(utils.MiddlewareUserUuidKey, user.Uuid)
	c.Set(utils.MiddlewareUsernameKey, user.Username)
	c.Set(utils.MiddlewareUserRoleKey, user.IsAdmin)
	c.Next()
}

func (h TokenHandler) AdminAuthorize(c *gin.Context) {

	isAdmin, _ := c.Get(utils.MiddlewareUserRoleKey)
	//if !ok {
	//	utils.HandleError(c, http.StatusInternalServerError, errors.New("cannot get role in middleware"))
	//	return
	//}

	if !isAdmin.(bool) {
		utils.HandleError(c, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	c.Next()
}
