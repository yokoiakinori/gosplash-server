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