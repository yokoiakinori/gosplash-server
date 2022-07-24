package service

import (
	"errors"

	"gosplash-server/app/model"

	"github.com/gin-contrib/sessions"
    "github.com/gin-contrib/sessions/cookie"
)

type UserService struct {
	
}

func (UserService) Register(user *model.User) error {
	_, err := DbEngine.Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (UserService) Login(c *gin.Context) error {
	user := model.User{}
	email := c.PostForm("email")
	password := c.PostForm("password")

	_, err := DbEngine.Where("email = ?", email).Get(&user)
	if user.Password != password {
		c.String(http.StatusBadRequest, "パスワードが一致しません。")
		return
	}
	
	if err != nil {
		return err
	}

	session := sessions.Default(c)
	session.Set("loginUser", email)
	session.Save()
	return nil
}