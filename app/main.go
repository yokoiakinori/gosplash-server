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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/joho/godotenv"

	"gosplash-server/app/controller"
	"gosplash-server/app/middleware"
	"gosplash-server/app/helper"
)

func loadEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 
}

func newS3() (*s3.S3, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	loadEnv()

	ak := os.Getenv("AWS_ACCESS_KEY_ID")
	sk := os.Getenv("AWS_SECRET_ACCESS_KEY")
	cfg := aws.Config{
		Credentials: credentials.NewStaticCredentials(ak, sk, ""),
		Region: aws.String("ap-northeast-1"),
		Endpoint: aws.String("http://minio:9001"),
		S3ForcePathStyle: aws.Bool(true), // s3のパスとminioのパスの形式が違うためこの1行が必要
	}
	return s3.New(s, &cfg), nil
}

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