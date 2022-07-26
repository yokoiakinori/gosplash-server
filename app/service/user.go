package service

import (
	"net/http"

	"gosplash-server/app/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

type UserService struct {
	
}

func (UserService) Register(c *gin.Context) {
	user := model.User{}
	if c.PostForm("password") != c.PostForm("password_confirmation") {
		c.String(http.StatusUnprocessableEntity, "パスワードと確認用パスワードが一致しません。")
		return
	}

	_, err := DbEngine.Insert(user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
	return
}

func (UserService) Login(c *gin.Context) {
	user := model.User{}
	email := c.PostForm("email")
	password := c.PostForm("password")

	_, err := DbEngine.Where("email = ?", email).Get(&user)
	if user.Password != password {
		c.String(http.StatusBadRequest, "パスワードが一致しません。")
		return
	}
	
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	session := sessions.Default(c)
	session.Set("loginUser", email)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"status": "ログイン完了",
	})
	return
}

func (UserService) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"status": "ログアウトしました。",
	})
	return
}

func (UserService) GetMyInfo(c *gin.Context) {
	user := model.User{}
	email, _ := c.Get("loginUser")
	_, err := DbEngine.Where("email = ?", email).Get(&user)
	
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	session := sessions.Default(c)
	session.Set("loginUser", email)
	session.Save()

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}