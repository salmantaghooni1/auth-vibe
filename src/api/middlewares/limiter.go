package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"

	"github.com/salmantaghooni/golang-car-web-api/api/helper"
)

func LimitByRequest() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(10, nil)
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, helper.OtpLimiterError, err))
			return
		} else {
			ctx.Next()
		}
	}
}
