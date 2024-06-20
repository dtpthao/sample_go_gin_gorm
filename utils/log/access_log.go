package log

import (
	"time"

	"github.com/gin-gonic/gin"
)

func AccessLog(c *gin.Context) {
	start := time.Now()
	path := c.Request.URL.Path
	raw := c.Request.URL.RawQuery
	if raw != "" {
		path = path + "?" + raw
	}
	Log().
		Str("IP", c.ClientIP()).
		Str("Method", c.Request.Method).
		Str("Path", path).
		Str("Protocol", c.Request.Proto).
		Str("Agent", c.Request.UserAgent()).
		Msg("request")
	c.Next()
	now := time.Now()
	Log().
		Str("IP", c.ClientIP()).
		Str("Method", c.Request.Method).
		Str("Path", path).
		Int("Response code", c.Writer.Status()).
		Str("Latency", now.Sub(start).String()).
		Msg("response")
}
