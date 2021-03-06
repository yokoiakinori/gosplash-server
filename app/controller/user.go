package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"gosplash-server/app/model"
	"gosplash-server/app/service"
)

func Register(c *gin.Context) {
	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	if c.PostForm("password") != c.PostForm("password_confirmation") {
		c.String(http.StatusUnprocessableEntity, "パスワードと確認用パスワードが一致しません。")
		return
	}

	userService := service.UserService{}
	err = userService.Register(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
}