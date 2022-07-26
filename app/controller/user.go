package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gosplash-server/app/model"
	"gosplash-server/app/service"
)

type User struct {

}

func (User) Register(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	userService := service.UserService{}
	userService.Register(c)
}

func (User) Login(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	userService := service.UserService{}
	userService.Login(c)
}

func (User) Logout(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	userService := service.UserService{}
	userService.Logout(c)
}

func (User) GetMyInfo(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	userService := service.UserService{}
	userService.GetMyInfo(c)
}

func (User) UpdateProfile(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	userService := service.UserService{}
	userService.UpdateProfile(c)
}