package s3uploader

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Hander s3 upload handler
type Hander struct {
	session *session.Session
	bucket  string
}

// Init s3 helper
func Init(region, bucket string) (Hander, error) {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials("AKIAIRDKW6IACJ6IBGCQ", "vqdEnq9I/Mrd4vwGrnd2oGwEI2rA7mLcn7BMPw4f", ""),
	})
	if err != nil {
		return Hander{}, err
	}
	return Hander{
		session: session,
		bucket:  bucket,
	}, nil
}

// NewUploader create new uploader
func (handler Hander) NewUploader(filePath, fileName string) *S3UploadHandler {
	return &S3UploadHandler{
		s3Uploader: s3manager.NewUploader(handler.session),
		fileName:   fileName,
		filePath:   filePath,
		bucket:     handler.bucket,
	}
}

// NewPuter create new puter
func (handler Hander) NewPuter(filePath, fileName string) S3PutObjectHandler {
	return S3PutObjectHandler{
		s3Client: s3.New(handler.session),
		fileName: fileName,
		filePath: filePath,
		bucket:   handler.bucket,
	}
}
