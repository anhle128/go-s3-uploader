package s3uploader

import (
	"bytes"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// S3UploadHandler s3 upload handler
type S3UploadHandler struct {
	s3Uploader  *s3manager.Uploader
	bucket      string
	filePath    string
	fileName    string
	responseURL string
}

// Write implement io.Writer
func (handler *S3UploadHandler) Write(p []byte) (n int, err error) {
	err = handler.Upload(p)
	if err != nil {
		return -1, err
	}
	return len(p), nil
}

// Upload load to s3
func (handler *S3UploadHandler) Upload(data []byte) error {
	upParams := &s3manager.UploadInput{
		Bucket:             aws.String(handler.bucket),
		Key:                aws.String(fmt.Sprintf("%s/%s", handler.filePath, handler.fileName)),
		Body:               bytes.NewReader(data),
		ContentDisposition: aws.String("attachment"),
		ACL:                aws.String(s3.ObjectCannedACLPublicRead),
	}
	// response, err := handler.s3Uploader.Upload(upParams, func(u *s3manager.Uploader) {
	// 	u.PartSize = 5 * 1024 * 1024 // 5MB part size
	// 	u.LeavePartsOnError = true   // Don't delete the parts if the upload fails.
	// })
	response, err := handler.s3Uploader.Upload(upParams)
	if err != nil {
		return err
	}
	handler.responseURL = response.Location
	return nil
}

// URL get url after upload success
func (handler *S3UploadHandler) URL() string {
	return handler.responseURL
}
