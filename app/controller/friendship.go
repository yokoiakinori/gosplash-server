package controller

import (
	"github.com/gin-gonic/gin"

	"gosplash-server/app/service"
)

type Friendship struct {

}

func (Friendship) Follow(c *gin.Context) {
	friendshipService := service.FriendshipService{}
	friendshipService.Follow(c)
}

func (Friendship) Unfollow(c *gin.Context) {
	friendshipService := service.FriendshipService{}
	friendshipService.Unfollow(c)
}