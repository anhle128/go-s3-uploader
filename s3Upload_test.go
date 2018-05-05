package s3uploader_test

import (
	"testing"

	s3uploader "github.com/anhle128/go-s3-uploader"
)

func TestUploader(t *testing.T) {
	s3handler, err := s3uploader.Init("ap-southeast-1", "football-x-uploaded")
	if err != nil {
		t.Error(err)
	}

	s3Uploader := s3handler.NewUploader("test", "motconvit")
	_, err = s3Uploader.Upload([]byte("mot con vit xue ra hai cai canh"))
	if err != nil {
		t.Error(err)
	}
}
