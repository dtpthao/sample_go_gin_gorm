package delivery

import (
	"errors"
	"github.com/gin-gonic/gin"
	"glintecoTask/entity"
	"glintecoTask/utils"
	"net/http"
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

	user, err := h.uuc.GetDetails(token.Username)
	if err != nil {
		utils.HandleError(c, http.StatusBadRequest, err)
		return
	}

	c.Set("username", user.Username)
	c.Set("isAdmin", user.IsAdmin)
	c.Next()
}

func (h TokenHandler) AdminAuthorize(c *gin.Context) {

	isAdmin, ok := c.Get("isAdmin")
	if !ok {
		utils.HandleError(c, http.StatusInternalServerError, errors.New("cannot get role in middleware"))
		return
	}

	if !isAdmin.(bool) {
		utils.HandleError(c, http.StatusUnauthorized, errors.New("unauthorized"))
		return
	}

	c.Next()
}
