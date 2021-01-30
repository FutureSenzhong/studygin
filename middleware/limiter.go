package middleware

import (
	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth_gin"
	"github.com/gin-gonic/gin"
)

func MyLimit() gin.HandlerFunc {

	l := tollbooth.NewLimiter(1, nil)
	l.SetMessage("服务繁忙，请稍后再试...")
	return tollbooth_gin.LimitHandler(l)
}
