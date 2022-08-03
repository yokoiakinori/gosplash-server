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
	err := c.Bind(&friendship)
	authorizerId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	user := model.User{}
	email, _ := c.Get("loginUser")
	_, err = DbEngine.Where("email = ?", email).Get(&user)

	if user.Id == authorizerId {
		c.String(http.StatusUnprocessableEntity, "自分自身はフォローできません。")
		return
	}

	friendship.ApplicantId = user.Id
	friendship.AuthorizerId = authorizerId

	_, err = DbEngine.Insert(friendship)
	if err != nil {
		panic(err)
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
	return
}

func (FriendshipService) Unfollow(c *gin.Context) {
	friendship := model.Friendship{}
	authorizerId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	user := model.User{}
	email, _ := c.Get("loginUser")
	_, err := DbEngine.Where("email = ?", email).Get(&user)

	_, err = DbEngine.Where("applicant_id = ?", user.Id).And("authorizer_id = ?", authorizerId).Get(&friendship)

	DbEngine.Delete(friendship)
	if err != nil {
		c.String(http.StatusInternalServerError, "Server Error")
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "ok",
	})
	return
}