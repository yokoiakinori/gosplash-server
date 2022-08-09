package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	_ "github.com/go-sql-driver/mysql"

	"gosplash-server/app/controller"
	"gosplash-server/app/middleware"
)

func main() {
	router := gin.Default()
	
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Use(middleware.RecordUaAndTime)
	router.MaxMultipartMemory = 8 << 20

	v1 := router.Group("/v1")
	{
		user := v1.Group("/users")
		{
			userController := controller.User{}
			user.POST("/register", userController.Register)
			user.POST("/login", userController.Login)
			user.POST("/logout", userController.Logout)
		}

		post := v1.Group("/posts")
		{
			postController := controller.Post{}
			post.GET("/", postController.GetAllPost)
			post.GET("/:id", postController.GetPost)
		}

		// ユーザー認証必要なルート
		auth := v1.Group("")
		auth.Use(middleware.LoginCheck())
		{
			user := auth.Group("/users")
			{
				userController := controller.User{}
				me := user.Group("/me")
				{
					me.GET("/", userController.GetMyInfo)
					me.PUT("/", userController.UpdateProfile)
					me.POST("/icon", userController.UpdateIcon)
				}
			}

			friendship := auth.Group("/friendships")
			{
				friendshipController := controller.Friendship{}
				friendship.POST("/:id", friendshipController.Follow)
				friendship.DELETE("/:id", friendshipController.Unfollow)
				friendship.GET("/", friendshipController.GetFollowers)
			}

			post := auth.Group("/posts")
			{
				postController := controller.Post{}
				post.POST("/", postController.Store)
				post.PUT("/:id", postController.Update)
				post.DELETE("/:id", postController.Delete)
			}

			like := auth.Group("/likes")
			{
				postController := controller.Post{}
				like.POST("/:id", postController.Like)
				like.DELETE("/:id", postController.Unlike)
			}

			comment := auth.Group("/comments")
			{
				postController := controller.Post{}
				comment.POST("/:id", postController.StoreComment)
			}
		}
	}

	router.Run(":8000")
}