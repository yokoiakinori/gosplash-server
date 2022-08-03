package service

import (
	"net/http"
	"strconv"

	"gosplash-server/app/model"

	"github.com/gin-gonic/gin"
)

type FriendshipService struct {
	
}

func (FriendshipService) Follow(c *gin.Context) {
	friendship := model.Friendship{}
	authorizerId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	user := model.User{}
	email, _ := c.Get("loginUser")
	_, err := DbEngine.Where("email = ?", email).Get(&user)

	if user.Id == authorizerId {
		c.String(http.StatusUnprocessableEntity, "自分自身はフォローできません。")
		return
	}

	_, err = DbEngine.Insert(friendship)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
	return
}