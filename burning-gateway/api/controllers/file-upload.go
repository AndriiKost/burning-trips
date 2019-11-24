package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/andriikost/burning-gateway/api/responses"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// type FileUploadRequest struct {
// 	fileName   string
// 	filePrefix string
// }

// func (fileUpload *FileUploadRequest) Prepare() {
// 	fileUpload.fileName = html.EscapeString(strings.TrimSpace(fileUpload.fileName))
// 	fileUpload.filePrefix = html.EscapeString(strings.TrimSpace(fileUpload.filePrefix))
// }

// func (fileUpload *FileUploadRequest) Validate() error {

// 	if fileUpload.fileName == "" {
// 		return errors.New("Required fileName")
// 	}
// 	if fileUpload.filePrefix == "" {
// 		return errors.New("Required filePrefix")
// 	}
// 	return nil
// }

func (server *Server) UploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("File Upload Endpoint Hit")

	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	awsSecret := os.Getenv("AWS_SECRET")
	awsUploadBucket := os.Getenv("AWS_UPLOAD_BUCKET")
	token := "FwoGZXIvYXdzECMaDKECLcU3A+p3rp784SJqoM5eob/wNh2Ri2Dd5L3XFOaCfBescWuyERnBrUl/NBaszBW2jtPwMEff5Z/3B7HN++WWGGpJT7tjAURzokeb3fqJ/zk0WEFBuRtKGzcdvEKmPXYh54dgEsY78w0+rsq68J2spnPLS4fqNSiA8OruBTIoS3ThijFfGI89c3Km3kBtwW1O7Hq4ofTplNfQUcImRl+Kb342EGSXKQ=="

	creds := credentials.NewStaticCredentials(awsAccessKey, awsSecret, token)

	cfg := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)
	svc := s3.New(session.New(), cfg)

	svc.

		// Parse our multipart form, 10 << 20 specifies a maximum
		// upload of 10 MB files.
		r.ParseMultipartForm(10 << 20)
	// FormFile returns the first file for the given key `myFile`
	// it also returns the FileHeader so we can get the Filename,
	// the Header and the size of the file
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}

	// fileSize := handler.Size
	filename := handler.Filename
	defer file.Close()

	// new way

	sess := session.Must(session.NewSession())

	uploader := s3manager.NewUploader(sess)

	f, err := os.Open(filename)
	if err != nil {
		// return fmt.Errorf("failed to open file %q, %v", filename, err)
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(awsUploadBucket),
		Key:    aws.String(awsAccessKey),
		Body:   f,
	})
	if err != nil {
		// return fmt.Errorf("failed to upload file, %v", err)
	}
	fmt.Printf("file uploaded to, %s\n", result.Location)

	// end

	// buffer := make([]byte, handler.Size)
	// file.Read(buffer)
	// fileBytes := bytes.NewReader(buffer)
	// fileType := http.DetectContentType(buffer)

	// path := "/stops/" + filename

	// params := &s3.PutObjectInput{
	// 	Bucket:        aws.String(os.Getenv("AWS_UPLOAD_BUCKET")),
	// 	Key:           aws.String(path),
	// 	Body:          fileBytes,
	// 	ContentLength: aws.Int64(fileSize),
	// 	ContentType:   aws.String(fileType),
	// }

	// resp, err := svc.PutObject(params)
	// if err != nil {
	// 	// handle error
	// }
	// fmt.Printf("response %s", awsutil.StringValue(resp))

	responses.JSON(w, http.StatusOK, result.Location)

	// START AWS
	// var bucket, key string
	// var timeout time.Duration

	// flag.StringVar(&bucket, "b", os.Getenv("AWS_UPLOAD_BUCKET"), "Bucket name.")
	// flag.StringVar(&key, "k", os.Getenv("AWS_ACCESS_KEY"), "Object key name.")
	// flag.DurationVar(&timeout, "d", 10000, "Upload timeout.")
	// flag.Parse()

	// // All clients require a Session. The Session provides the client with
	// // shared configuration such as region, endpoint, and credentials. A
	// // Session should be shared where possible to take advantage of
	// // configuration and credential caching. See the session package for
	// // more information.
	// sess := session.Must(session.NewSession(&aws.Config{
	// 	Region: aws.String("us-east-1"),
	// }))

	// // Create a new instance of the service's client with a Session.
	// // Optional aws.Config values can also be provided as variadic arguments
	// // to the New function. This option allows you to provide service
	// // specific configuration.
	// svc := s3.New(sess)

	// // Create a context with a timeout that will abort the upload if it takes
	// // more than the passed in timeout.
	// ctx := context.Background()
	// var cancelFn func()
	// if timeout > 0 {
	// 	ctx, cancelFn = context.WithTimeout(ctx, timeout)
	// }
	// // Ensure the context is canceled to prevent leaking.
	// // See context package for more information, https://golang.org/pkg/context/
	// if cancelFn != nil {
	// 	defer cancelFn()
	// }

	// // Uploads the object to S3. The Context will interrupt the request if the
	// // timeout expires.
	// putObjectOutput, err := svc.PutObjectWithContext(ctx, &s3.PutObjectInput{
	// 	Bucket: aws.String(bucket),
	// 	Key:    aws.String(key),
	// 	Body:   os.Stdin,
	// })
	// fmt.Println(putObjectOutput)
	// if err != nil {
	// 	if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
	// 		// If the SDK can determine the request or retry delay was canceled
	// 		// by a context the CanceledErrorCode error code will be returned.
	// 		fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
	// 	} else {
	// 		fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
	// 	}
	// 	os.Exit(1)
	// }

	// fmt.Printf("successfully uploaded file to %s/%s\n", bucket, key)
}
