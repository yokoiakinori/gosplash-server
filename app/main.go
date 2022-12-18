package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	_ "github.com/go-sql-driver/mysql"

	"gosplash-server/app/controller"
	"gosplash-server/app/middleware"
)

func main() {
	router := gin.Default()

	// CORSの設定
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string{
			"GET",
			"POST",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	}))
	
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	v1 := router.Group("/v1")
	{
		userController := controller.User{}
		v1.POST("/users/register", userController.Register)
		v1.POST("/users/login", userController.Login)
		v1.POST("/users/logout", userController.Logout)

		postController := controller.Post{}
		v1.GET("/posts", postController.GetAllPost)
		v1.GET("/posts/:id", postController.GetPost)

		// ユーザー認証必要なルート
		auth := v1.Group("")
		auth.Use(middleware.LoginCheck())
		{
			userController := controller.User{}
			auth.GET("/users/me", userController.GetMyInfo)
			auth.PUT("/users/me", userController.UpdateProfile)
			auth.POST("/users/me/icon", userController.UpdateIcon)

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
				post.PATCH("/:id", postController.Update)
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
				comment.PATCH("/:id", postController.UpdateComment)
				comment.DELETE("/:id", postController.DeleteComment)
			}

			collection := auth.Group("/collections")
			{
				postController := controller.Post{}
				collection.POST("/:id", postController.StoreCollection)
				collection.DELETE("/:id", postController.DeleteCollection)
				collection.GET("/", postController.GetCollections)
			}
		}
	}

	router.Run(":8000")
}