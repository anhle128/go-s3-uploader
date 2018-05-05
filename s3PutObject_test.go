package s3uploader_test

import (
	"testing"

	s3uploader "github.com/anhle128/go-s3-uploader"
)

func TestPuter(t *testing.T) {
	s3handler, err := s3uploader.Init("ap-southeast-1", "football-x-uploaded")
	if err != nil {
		t.Error(err)
	}

	s3PutObject := s3handler.NewPuter("test", "motconvit")
	_, err = s3PutObject.PutObject([]byte("mot con vit xue ra hai cai canh"))
	if err != nil {
		t.Error(err)
	}
}
