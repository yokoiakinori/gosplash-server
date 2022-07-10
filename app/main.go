package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"

	"app/controller"
	"app/middleware"
)

func main() {
	router := gin.Default()
	
	router.Use(middleware.RecordUaAndTime)

	bookRouter := router.Group("/book")
	{
		v1 := bookRouter.Group("/v1")
		{
			v1.POST("/", controller.BookAdd)
			v1.GET("/", controller.BookList)
			v1.PUT("/", controller.BookUpdate)
			v1.DELETE("/", controller.BookDelete)
		}
	}

	router.Run(":8000")
}