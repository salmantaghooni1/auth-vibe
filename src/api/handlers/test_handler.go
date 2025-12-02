package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/salmantaghooni/golang-car-web-api/api/helper"
)

type header struct {
	UserId  string
	Browser string
}

type PersonData struct {
	FirstName    string `json:"first_name" binding:"required,alpha,min=4,max=255"`
	LastName     string `json:"last_name" binding:"required,alpha,min=4,max=255"`
	MobileNumber string `json:"mobile_number" binding:"required,mobile,min=11,max=11"`
}

type TestHandler struct {
}

func NewTesthHandler() *TestHandler {
	return &TestHandler{}

}

func (h *TestHandler) Test(c *gin.Context) {

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"resault": "test",
	}, true, 0))

}

func (h *TestHandler) Users(c *gin.Context) {

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"resault": "users",
	}, true, 0))
}

// UserByID godoc
// @Summery UserByID
// @Description UserByID
// @Tags Test
// @Accept json
// @Produce json
// @Param id path int true "user id"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/user/{id} [get]
func (h *TestHandler) UserByID(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"type":    "userbyid",
		"resault": id,
	}, true, 0))
}

func (h *TestHandler) UserByUsername(c *gin.Context) {
	userName := c.Param("username")

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"type":    "username",
		"resault": userName,
	}, true, 0))
}

func (h *TestHandler) AccountByID(c *gin.Context) {
	id := c.Param("id")

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"type":    "account",
		"resault": id,
	}, true, 0))
}

func (h *TestHandler) AddUser(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"type": "add user",
	}, true, 0))
}

func (h *TestHandler) HeaderBinder(c *gin.Context) {
	userId := c.GetHeader("UserId")

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"resault": "header binder1",
		"userId":  userId,
	}, true, 0))
}

func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	c.BindHeader(&header)

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"resault": "header binder2",
		"userId":  header,
	}, true, 0))
}

func (h *TestHandler) QueryBinder1(c *gin.Context) {
	// ides := c.QueryArray("id")
	id := c.Query("id")
	name := c.Query("name")

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"resault": "QueryBinder1",
		"userId":  id,
		"name":    name,
	}, true, 0))
}

func (h *TestHandler) UriBinder1(c *gin.Context) {
	// ides := c.QueryArray("id")
	id := c.Param("id")
	name := c.Param("name")

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"resault": "UriBinder1s",
		"userId":  id,
		"name":    name,
	}, true, 0))
}

// BodyBinder godoc
// @Summery BodyBinder
// @Description BodyBinder
// @Tags Test
// @Accept json
// @Produce json
// @Param person body PersonData true "person data"
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/test/body [post]
// @Security AuthBearer
func (h *TestHandler) BodyBinder(c *gin.Context) {
	p := PersonData{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"validation": err.Error(),
		})
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"resault": "BodyBinder",
	}, true, 0))
}

func (h *TestHandler) FromBinder(c *gin.Context) {
	p := PersonData{}
	err := c.ShouldBind(&p)
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"resault": "FromBinder",
	}, true, 0))
}

func (h *TestHandler) FileBinder(c *gin.Context) {

	file, _ := c.FormFile("file")
	err := c.SaveUploadedFile(file, "file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, helper.GenerateBaseResponse(gin.H{
		"resault": file.Filename,
	}, true, 0))

}
