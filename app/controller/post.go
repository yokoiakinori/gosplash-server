package controller

import (
	"github.com/gin-gonic/gin"

	"gosplash-server/app/service"
)

type Post struct {

}

func (Post) Store(c *gin.Context) {
	postService := service.PostService{}
	postService.Store(c)
}

func (Post) Update(c *gin.Context) {
	postService := service.PostService{}
	postService.Update(c)
}

func (Post) Delete(c *gin.Context) {
	postService := service.PostService{}
	postService.Delete(c)
}

func (Post) GetAllPost(c *gin.Context) {
	postService := service.PostService{}
	postService.GetAllPost(c)
}

func (Post) GetPost(c *gin.Context) {
	postService := service.PostService{}
	postService.GetPost(c)
}

func (Post) Like(c *gin.Context) {
	postService := service.PostService{}
	postService.Like(c)
}

func (Post) Unlike(c *gin.Context) {
	postService := service.PostService{}
	postService.Unlike(c)
}

func (Post) StoreComment(c *gin.Context) {
	postService := service.PostService{}
	postService.StoreComment(c)
}

func (Post) UpdateComment(c *gin.Context) {
	postService := service.PostService{}
	postService.UpdateComment(c)
}

func (Post) DeleteComment(c *gin.Context) {
	postService := service.PostService{}
	postService.DeleteComment(c)
}

func (Post) StoreCollection(c *gin.Context) {
	postService := service.PostService{}
	postService.StoreCollection(c)
}

func (Post) DeleteCollection(c *gin.Context) {
	postService := service.PostService{}
	postService.DeleteCollection(c)
}

func (Post) GetCollections(c *gin.Context) {
	postService := service.PostService{}
	postService.GetCollections(c)
}