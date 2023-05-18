package handlers

import (
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

const (
	maxPartSize        = int64(5 * 1024 * 1024)
	maxRetries         = 3
	awsAccessKeyID     = "AKIATFDGZ4RATBZE2MXH"
	awsSecretAccessKey = "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"
	awsBucketRegion    = "us-east-1"
	awsBucketName      = "taxtrackingbucket"
)

var myBucket = "taxtrackingbucket"
var accessKey = "AKIATFDGZ4RATBZE2MXH"
var accessSecret = "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"

func FileUploadHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, files := range r.MultipartForm.File {
		for _, file := range files {
			// Open the file
			src, err := file.Open()
			filename := file.Filename
			if err != nil {
				fmt.Println(err)
				return
			}
			defer src.Close()
			image := UploadFileToS3(src, filename)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(image))
			fmt.Println(image)
		}
	}
}

func UploadFileToS3(filee multipart.File, name string) string {
	var awsConfig *aws.Config
	if accessKey == "" || accessSecret == "" {
		//load default credentials
		awsConfig = &aws.Config{
			Region: aws.String("us-east-1"),
		}
	} else {
		awsConfig = &aws.Config{
			Region:      aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(accessKey, accessSecret, ""),
		}
	}

	// The session the S3 Uploader will use
	sess := session.Must(session.NewSession(awsConfig))

	// Create an uploader with the session and default options
	//uploader := s3manager.NewUploader(sess)

	// Create an uploader with the session and custom options
	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024 // The minimum/default allowed part size is 5MB
		u.Concurrency = 2            // default is 5
	})

	// Upload the file to S3.
	result, _ := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String("products/" + name),
		Body:   filee,
		ACL:    aws.String("public-read"),
	})

	return result.Location
}
