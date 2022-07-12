package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"gosplash-server/app/controller"
	"gosplash-server/app/middleware"
)

func main() {
	router := gin.Default()
	
	router.Use(middleware.RecordUaAndTime)

	v1 := router.Group("/v1")
	{
		user := v1.Group("/users")
		{
			user.POST("/register", controller.Register)
			user.POST("/login", controller.Login)
		}
	}

	router.Run(":8000")
}