package service

import (
	"net/http"
	"os"

	"gosplash-server/app/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"gosplash-server/app/helper"
	"gosplash-server/app/setup"
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

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
	return
}

func (UserService) UpdateProfile(c *gin.Context) {
	email, _ := c.Get("loginUser")

	user := model.User {
		Name: c.PostForm("name"),
		Description: c.PostForm("description"),
	}

	_, err := DbEngine.Where("email = ?", email).Update(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func (UserService) UpdateIcon(c *gin.Context) {
	fileHeader, _ := c.FormFile("file")
	fileName := fileHeader.Filename
	user := model.User{}
	email, _ := c.Get("loginUser")

	_, err := DbEngine.Where("email = ?", email).Get(&user)
	filePath, err := helper.MakeFilePath("icon", fileName)
	icon := model.Icon {
		Path: filePath,
		UserId: user.Id,
	}
	_, err = DbEngine.Insert(&icon)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	var bucket = os.Getenv("MINIO_DEFAULT_BUCKETS")
	s3Session, err := newS3()
	params := &s3.PutObjectInput {
		Bucket: aws.String(bucket),
		Key: filePath,
		Body: file,
	}
	_, err = s3Session.PutObject(params)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}