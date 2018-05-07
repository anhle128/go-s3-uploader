package s3uploader

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// S3PutObjectHandler s3 put object handler
type S3PutObjectHandler struct {
	s3Client *s3.S3
	bucket   string
	filePath string
	fileName string
}

// PutObject put object to s3
func (handler S3PutObjectHandler) PutObject(data []byte) (*s3.PutObjectOutput, error) {
	return handler.s3Client.PutObject(&s3.PutObjectInput{
		Bucket:             aws.String(handler.bucket),
		Key:                aws.String(fmt.Sprintf("%s/%s", handler.filePath, handler.fileName)),
		ACL:                aws.String(s3.ObjectCannedACLPublicRead),
		Body:               bytes.NewReader(data),
		ContentDisposition: aws.String("attachment"),
	})
}

// Write implement io.Writer
func (handler S3PutObjectHandler) Write(p []byte) (n int, err error) {
	_, err = handler.PutObject(p)
	if err != nil {
		return -1, err
	}
	return len(p), nil
}
