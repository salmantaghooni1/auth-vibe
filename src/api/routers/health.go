package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/salmantaghooni/golang-car-web-api/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHandler()
	r.GET("/", handler.Health)
}
