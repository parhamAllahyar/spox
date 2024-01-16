package s3

import (
	"asset/config"
	"asset/internal/core/port/driven/storage"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	s3Pkg "github.com/aws/aws-sdk-go/service/s3"
)

type s3 struct {
	config config.S3
}

func NewS3(config config.S3) storage.Storage {
	return &s3{
		config: config,
	}

}

func (s3 s3) newSession() (*session.Session, error) {

	return session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(s3.config.ACCESSKEY, s3.config.SECRETKEY, s3.config.Token),
		Region:      aws.String(s3.config.Region),
		Endpoint:    aws.String(s3.config.Endpoint),
	})

}

func (s3 s3) newSVCSession() (*s3Pkg.S3, error) {

	newSession, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(s3.config.ACCESSKEY, s3.config.SECRETKEY, s3.config.Token),
	})

	if err != nil {
		return nil, nil
	}

	svc := s3Pkg.New(newSession, &aws.Config{
		Region:   aws.String(""),
		Endpoint: aws.String("<ENDPOINT_URL>"),
	})

	return svc, nil

}
