package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"

	"github.com/salmantaghooni/golang-car-web-api/api/helper"
	"github.com/salmantaghooni/golang-car-web-api/config"
	"github.com/salmantaghooni/golang-car-web-api/pkg/limmiter"
)

func OtpLimiter(cfg *config.Config) gin.HandlerFunc {
	var limiter = limmiter.NewIPRateLimiter(rate.Every(cfg.OTP.Limmiter*time.Second), 1)
	return func(c *gin.Context) {
		limiter := limiter.GetLimiter(c.Request.RemoteAddr)
		if !limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, http.StatusTooManyRequests, nil))
			c.Abort()
		} else {
			c.Next()
		}
	}
}
