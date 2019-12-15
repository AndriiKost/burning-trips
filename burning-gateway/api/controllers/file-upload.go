package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gorilla/mux"
)

type PresignResp struct {
	Method, URL string
	Header      http.Header
}

func (server *Server) GetPresignedUploadUrl(w http.ResponseWriter, r *http.Request) {

	var u string
	var signedHeaders http.Header

	vars := mux.Vars(r)
	fmt.Println(vars)
	key := vars["object-name"]

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

	// For creating PutObject presigned URLs
	fmt.Println("Received request to presign PutObject for,", key)
	sdkReq, _ := s3Svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_UPLOAD_BUCKET")),
		Key:    aws.String(key),
	})
	u, signedHeaders, err = sdkReq.PresignRequest(15 * time.Minute)

	// Create the response back to the client with the information on the
	// presigned request and additional headers to include.
	if err := json.NewEncoder(w).Encode(PresignResp{
		URL:    u,
		Header: signedHeaders,
	}); err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode presign response, %v", err)
	}

}
