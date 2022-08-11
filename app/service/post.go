package service

import (
	"net/http"
	"strconv"
	"os"

	"gosplash-server/app/model"
	"gosplash-server/app/helper"
	"gosplash-server/app/setup"

	"github.com/gin-gonic/gin"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
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

func (PostService) Delete(c *gin.Context) {
	session := DbEngine.NewSession()
	defer session.Close()

	post := model.Post{}
	postId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	_, err := DbEngine.Where("id = ?", postId).Get(&post)
	filePath := post.Path
	DbEngine.Delete(post)
	if err != nil {
		session.Rollback()
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	err = DeleteFile(filePath)
	if err != nil {
		session.Rollback()
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func (PostService) GetAllPost(c *gin.Context) {
	posts := []model.PostListAndUser{}
	err := DbEngine.Table("post").
	Join("INNER", "user", "user.id = post.user_id").
	Join("INNER", "icon", "icon.user_id = post.user_id").
	Find(&posts)

	if err != nil {
		panic(err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": posts,
	})
	return
}

func (PostService) GetPost(c *gin.Context) {
	post := model.PostAndUser{}
	_, err := DbEngine.Table("post").
	Join("INNER", "user", "user.id = post.user_id").
	Join("INNER", "icon", "icon.user_id = post.user_id").
	Where("post.id = ?", c.Param("id")).
	Get(&post)

	updatePost := model.PostAndUser{
		ViewCount: post.ViewCount + 1,
	}
	_, err = DbEngine.Table("post").
	Where("id = ?", c.Param("id")).Update(&updatePost)

	if err != nil {
		panic(err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"data": post,
	})
	return
}

func (PostService) Like(c *gin.Context) {
	email, _ := c.Get("loginUser")

	user := model.User {}

	_, err := DbEngine.Where("email = ?", email).Get(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	postId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	like := model.Like{
		UserId: user.Id,
		PostId: postId,
	}
	_, err = DbEngine.Insert(&like)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func (PostService) Unlike(c *gin.Context) {
	email, _ := c.Get("loginUser")

	user := model.User {}

	_, err := DbEngine.Where("email = ?", email).Get(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	postId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	like := model.Like {}
	_, err = DbEngine.Where("user_id = ?", user.Id).And("post_id = ?", postId).Delete(&like)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func (PostService) StoreComment(c *gin.Context) {
	email, _ := c.Get("loginUser")

	user := model.User {}

	_, err := DbEngine.Where("email = ?", email).Get(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	postId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	comment := model.Comment{
		UserId: user.Id,
		PostId: postId,
		Content: c.PostForm("content"),
	}
	_, err = DbEngine.Insert(&comment)

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
	return
}

func (PostService) UpdateComment(c *gin.Context) {
	commentId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	comment := model.Comment{}
	comment.Content = c.PostForm("content")
	_, err := DbEngine.ID(commentId).Update(&comment)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func (PostService) DeleteComment(c *gin.Context) {
	commentId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	comment := model.Comment{}
	_, err := DbEngine.ID(commentId).Delete(&comment)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}

func (PostService) StoreCollection(c *gin.Context) {
	email, _ := c.Get("loginUser")

	user := model.User {}

	_, err := DbEngine.Where("email = ?", email).Get(&user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	postId, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	comment := model.Collection{
		UserId: user.Id,
		PostId: postId,
	}
	_, err = DbEngine.Insert(&comment)

	c.JSON(http.StatusCreated, gin.H{
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

func DeleteFile(filePath string) (error) {
	var bucket = os.Getenv("MINIO_DEFAULT_BUCKETS")
	s3Session, err := setup.NewS3()
	params := &s3.DeleteObjectInput {
		Bucket: aws.String(bucket),
		Key: aws.String(filePath),
	}
	_, err = s3Session.DeleteObject(params)
	return err
}