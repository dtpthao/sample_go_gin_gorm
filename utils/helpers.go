package utils

import (
	"github.com/gin-gonic/gin"
	"glintecoTask/utils/log"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

func HandleError(c *gin.Context, status int, err error) {
	log.Error().Err(err)
	// todo response message should be filtered
	c.AbortWithStatusJSON(status, ErrorResponse{Message: err.Error()})
}
