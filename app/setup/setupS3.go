package setup

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func NewS3() (*s3.S3, error) {
	s, err := session.NewSession()
	if err != nil {
		return nil, err
	}

	LoadEnv()

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