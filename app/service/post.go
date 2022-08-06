package service

import (
	"net/http"
	"strconv"

	"gosplash-server/app/model"
	"gosplash-server/app/helper"

	"github.com/gin-gonic/gin"
)

type PostService struct {
	
}

func (PostService) Store(c *gin.Context) {
	session := DbEngine.NewSession()
	defer session.Close()

	session.Begin()
	fileHeader, _ := c.FormFile("file")
	fileName := fileHeader.Filename
	user, err := GetUserInfo(c)

	post, err := InsertPostRecord(c, fileName, user)
	if err != nil {
		session.Rollback()
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	file, _ := fileHeader.Open()
	err = UploadFile(post.Path, file)
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

func (PostService) Update(c *gin.Context) {
	post := model.Post{
		Name: c.PostForm("name"),
		Description: c.PostForm("description"),
	}
	postId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	_, err := DbEngine.Where("id = ?", postId).Update(&post)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func InsertPostRecord(c *gin.Context, fileName string, user model.User) (model.Post, error) {
	filePath, err := helper.MakeFilePath("post", fileName)

	post := model.Post {
		Name: c.PostForm("name"),
		Description: c.PostForm("description"),
		Path: filePath,
		UserId: user.Id,
	}
	_, err = DbEngine.Insert(&post)
	return post, err
}