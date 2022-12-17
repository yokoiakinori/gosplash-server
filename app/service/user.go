package service

import (
	"net/http"
	"os"
	"mime/multipart"

	"gosplash-server/app/model"
	"gosplash-server/app/helper"
	"gosplash-server/app/setup"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type UserService struct {
	
}

func (UserService) Register(input *model.Register) (error) {
	user := model.User{
		Name: input.Name,
		Email: input.Email,
		Password: input.Password,
	}

	_, err := DbEngine.Insert(&user)
	if err != nil {
		return err
	}

	return nil
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
	session := DbEngine.NewSession()
	defer session.Close()

	session.Begin()
	fileHeader, _ := c.FormFile("file")
	fileName := fileHeader.Filename
	user, err := GetUserInfo(c)

	err = deleteOldIconRecord(c, user.Id)
	if err != nil {
		session.Rollback()
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	icon, err := InsertIconRecord(c, fileName, user)
	if err != nil {
		session.Rollback()
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	file, _ := fileHeader.Open()
	err = UploadFile(icon.Path, file)
	if err != nil {
		session.Rollback()
		return
	}

	session.Commit()

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func GetUserInfo(c *gin.Context) (model.User, error){
	user := model.User{}
	email, _ := c.Get("loginUser")
	_, err := DbEngine.Where("email = ?", email).Get(&user)
	return user, err
}

func deleteOldIconRecord(c *gin.Context, userId int64) (error) {
	icon := model.Icon{}
	_, err := DbEngine.Where("user_id = ?", userId).Get(&icon)
	if icon.UserId == userId {
		DbEngine.Delete(icon)
	}
	return err
}

func InsertIconRecord(c *gin.Context, fileName string, user model.User) (model.Icon, error) {
	filePath, err := helper.MakeFilePath("icon", fileName)
	icon := model.Icon {
		Path: filePath,
		UserId: user.Id,
	}
	_, err = DbEngine.Insert(&icon)
	return icon, err
}

func UploadFile(filePath string, file multipart.File) (error) {
	var bucket = os.Getenv("MINIO_DEFAULT_BUCKETS")
	s3Session, err := setup.NewS3()
	params := &s3.PutObjectInput {
		Bucket: aws.String(bucket),
		Key: aws.String(filePath),
		Body: file,
	}
	_, err = s3Session.PutObject(params)
	return err
}