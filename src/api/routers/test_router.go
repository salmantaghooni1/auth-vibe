package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/salmantaghooni/golang-car-web-api/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTesthHandler()

	r.GET("/", h.Test)
	r.GET("/users", h.Users)
	r.GET("/user/:id", h.UserByID)
	r.GET("/user/username/:username", h.UserByUsername)
	r.GET("/user/:id/accounts", h.AccountByID)
	r.POST("/user/store", h.AddUser)

	//header binder
	r.POST("/header1", h.HeaderBinder)
	r.POST("/header2", h.HeaderBinder2)

	//query string binder
	r.POST("/querystring", h.QueryBinder1)

	//uri binder
	r.POST("/uri", h.UriBinder1)

	//body binder
	r.POST("/body", h.BodyBinder)

	//file binder
	r.POST("/file", h.FileBinder)
}
