package main

import (
	"os"
	"fmt"

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
	var filePath = "./androidparty.png"
	var bucket = os.Getenv("MINIO_DEFAULT_BUCKETS")
	var key = "sample"
	imageFile, imageErr := os.Open(filePath)
	if imageErr != nil {
		panic(imageErr)
	}
	defer imageFile.Close()
	c, err := newS3()
	if err != nil {
		panic(err)
	}
	params := &s3.PutObjectInput {
		Bucket: aws.String(bucket),
		Key: aws.String(key),
		Body: imageFile,
	}
	_, err = c.PutObject(params)
	if err != nil {
		panic(err)
	}

	router := gin.Default()
	
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	router.Use(middleware.RecordUaAndTime)

	v1 := router.Group("/v1")
	{
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