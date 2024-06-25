package log

import (
	"time"

	"github.com/gin-gonic/gin"
)

func AccessLog(ctx *gin.Context) {
	globalLog.Infof("Request: %v\n",
		ctx.ClientIP(),
		ctx.Request.Method,
		ctx.Request.URL.Path,
		ctx.Request.Proto,
		ctx.Request.UserAgent(),
		ctx.Request.Header.Get("Content-Type"),
	)
	start := time.Now()
	ctx.Next()
	globalLog.Infof("Response: %v\n",
		ctx.ClientIP(),
		ctx.Writer.Status(),
		ctx.Request.URL.Path,
		time.Since(start).String(),
		ctx.Writer.Header().Get("Content-Type"),
	)
}
