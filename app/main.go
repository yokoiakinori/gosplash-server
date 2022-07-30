package main

import (
	"os"
	"fmt"
	"net/http"
	// "log"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions"
	_ "github.com/go-sql-driver/mysql"

	"gosplash-server/app/controller"
	"gosplash-server/app/middleware"
	"gosplash-server/app/helper"
	"gosplash-server/app/setup"
)

func main() {
	// var bucket = os.Getenv("MINIO_DEFAULT_BUCKETS")
	// var key = "sample.png"

	router := gin.Default()
	
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Use(middleware.RecordUaAndTime)
	router.MaxMultipartMemory = 8 << 20

	v1 := router.Group("/v1")
	{
		image := v1.Group("/images")
		{
			image.POST("/upload", func(c *gin.Context){
				randomStr, err := helper.MakeFilePath("icon")
				if err != nil {
					c.String(http.StatusInternalServerError, "エラー")
				}
				c.JSON(http.StatusCreated, gin.H{
					"randomStr": randomStr,
				})
				// fileHeader, _ := c.FormFile("file")
				// log.Println(fileHeader.Filename)
				// file, _ := fileHeader.Open()

				// s3Session, err := newS3()
				// if err != nil {
				// 	log.Println(err)
				// }
				// params := &s3.PutObjectInput {
				// 	Bucket: aws.String(bucket),
				// 	Key: aws.String(key),
				// 	Body: file,
				// }
				// _, err = s3Session.PutObject(params)
				// if err != nil {
				// 	log.Println(err)
				// }
			})
		}
		user := v1.Group("/users")
		{
			userController := controller.User{}
			user.POST("/register", userController.Register)
			user.POST("/login", userController.Login)
			user.POST("/logout", userController.Logout)
		}

		// ユーザー認証必要なルート
		auth := v1.Group("")
		auth.Use(middleware.LoginCheck())
		{
			user := auth.Group("/users")
			{
				userController := controller.User{}
				user.GET("/me", userController.GetMyInfo)
				user.PUT("/:id", userController.UpdateProfile)
			}
		}
	}

	router.Run(":8000")
}