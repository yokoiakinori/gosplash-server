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