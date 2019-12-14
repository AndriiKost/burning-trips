package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/andriikost/burning-gateway/api/responses"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (server *Server) GetPresignedUploadUrl(w http.ResponseWriter, r *http.Request) {

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	awsSecret := os.Getenv("AWS_SECRET")
	awsUploadBucket := os.Getenv("AWS_UPLOAD_BUCKET")

	// token := "FwoGZXIvYXdzECMaDKECLcU3A+p3rp784SJqoM5eob/wNh2Ri2Dd5L3XFOaCfBescWuyERnBrUl/NBaszBW2jtPwMEff5Z/3B7HN++WWGGpJT7tjAURzokeb3fqJ/zk0WEFBuRtKGzcdvEKmPXYh54dgEsY78w0+rsq68J2spnPLS4fqNSiA8OruBTIoS3ThijFfGI89c3Km3kBtwW1O7Hq4ofTplNfQUcImRl+Kb342EGSXKQ=="

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, "")

	cfg := aws.NewConfig().WithRegion("us-east-2").WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(awsUploadBucket),
		Key:    aws.String(awsAccessKey),
	})

	url, err := req.Presign(3000 * time.Second)
	if err != nil {
		panic(err)
	}

	fmt.Println(url)

	responses.JSON(w, http.StatusOK, url)
}
