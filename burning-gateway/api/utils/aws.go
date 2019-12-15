package aws

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func getService() *s3.S3 {
	creds := credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY"), os.Getenv("AWS_SECRET"), "")

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(os.Getenv("AWS_PRIMARY_REGION")),
		Credentials: creds,
	})

	if err != nil {
		fmt.Printf("failed to create a session for aws")
	}

	s3Svc := s3.New(sess, &aws.Config{
		Region: aws.String(os.Getenv("AWS_PRIMARY_REGION")),
	})

	return s3Svc
}
