package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/andriikost/burning-gateway/api/responses"
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

type ObjectUploadResponse struct {
	ObjectUrl string `json:"objectUrl"`
	Success   bool   `json:"success"`
	Message   string `json:"message"`
}

const stopImagePrefix = "stops/"

func (server *Server) GetPresignedUploadUrl(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["object-name"]

	s3Svc := awsService()

	// For creating PutObject presigned URLs
	fmt.Println("Received request to presign PutObject for,", key)
	sdkReq, _ := s3Svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("AWS_UPLOAD_BUCKET")),
		Key:    aws.String(key),
	})
	url, signedHeaders, err := sdkReq.PresignRequest(15 * time.Minute)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to presign request, %v", err)
	}

	// Create the response back to the client with the information on the
	// presigned request and additional headers to include.
	if err := json.NewEncoder(w).Encode(PresignResp{
		URL:    url,
		Header: signedHeaders,
	}); err != nil {
		fmt.Fprintf(os.Stderr, "failed to encode presign response, %v", err)
	}

}

func awsService() *s3.S3 {
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

func (server *Server) UploadFile(w http.ResponseWriter, r *http.Request) {
	maxSize := int64(10240000)
	vars := mux.Vars(r)
	reqFileName := stopImagePrefix + vars["object-name"]

	err := r.ParseMultipartForm(maxSize)
	if err != nil {
		log.Println(err)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		return
	}

	defer file.Close()

	objectUrl, err := UploadFileToS3(file, fileHeader, reqFileName)

	if err != nil {
		log.Println(err)
		return
	}

	uploadMessage := "Image uploaded successfully: " + objectUrl
	fmt.Println(uploadMessage)

	response := ObjectUploadResponse{objectUrl, true, uploadMessage}

	responses.JSON(w, http.StatusOK, response)
}

func UploadFileToS3(file multipart.File, fileHeader *multipart.FileHeader, fileName string) (string, error) {

	// get the file size and read
	// the file content into a buffer
	size := fileHeader.Size
	buffer := make([]byte, size)
	file.Read(buffer)

	s3Svc := awsService()

	// create a unique file name for the file
	t := time.Now()
	tempFileName := fileName + t.Format("20060102150405") + filepath.Ext(fileHeader.Filename)

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	_, err := s3Svc.PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(os.Getenv("AWS_UPLOAD_BUCKET")),
		Key:                  aws.String(tempFileName),
		ACL:                  aws.String("public-read"), // could be private if you want it to be access by only authorized users
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	if err != nil {
		return "", err
	}

	fileUrl := "https://" + os.Getenv("AWS_UPLOAD_BUCKET") + ".s3." + os.Getenv("AWS_PRIMARY_REGION") + ".amazonaws.com/" + tempFileName

	return fileUrl, err
}
