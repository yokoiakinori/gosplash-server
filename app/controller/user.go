package controller

import (
	"net/http"
	"log"

	"github.com/gin-gonic/gin"

	"gosplash-server/app/model"
	"gosplash-server/app/service"
)

type User struct {

}

func (User) Register(c *gin.Context) {
	input := model.Register{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	if c.PostForm("password") != c.PostForm("password_confirmation") {
		c.JSON(http.StatusUnprocessableEntity, "パスワードと確認用パスワードが一致しません。")
		return
	}

	userService := service.UserService{}
	err = userService.Register(&input)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "会員登録をしました。",
	})
}

func (User) Login(c *gin.Context) {
	input := model.Login{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	userService := service.UserService{}
	userService.Login(&input, c)

	c.JSON(http.StatusOK, gin.H{
		"status": "ログイン完了",
	})
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

func (User) UpdateIcon(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	userService := service.UserService{}
	userService.UpdateIcon(c)
}