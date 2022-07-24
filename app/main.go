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

	v1 := router.Group("/v1")
	{
		user := v1.Group("/users")
		{
			userController = controller.User{}
			user.POST("/register", userController.Register)
			user.POST("/login", userController.Login)
			user.POST("/logout", userController.Logout)
			user.GET("/:id", userController.Logout)
		}
	}

	router.Run(":8000")
}