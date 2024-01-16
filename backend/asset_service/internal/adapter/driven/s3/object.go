package s3

import (
	"fmt"
	"io"
	"log"

	"time"

	"github.com/aws/aws-sdk-go/aws"
	s3Pkg "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// Remove Object
func (s3 s3) ObjectRemove() error {
	return nil
}

func (s3 s3) ObjectUpload(file io.Reader) error {

	// file, err := os.Open(path)
	// defer file.Close()

	session, err := s3.newSession()

	uploader := s3manager.NewUploader(session)
	fileData, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s3.config.Bucket),
		Key:    aws.String("filename"),
		Body:   file,
	})
	if err != nil {
		return err
	}

	fmt.Println(fileData)

	return nil
}

func (s3 s3) ObjectDownloadPerSign(file string) (string, error) {

	session, err := s3.newSVCSession()

	if err != err {
		log.Println(err)
	}

	req, _ := session.GetObjectRequest(&s3Pkg.GetObjectInput{
		Bucket: aws.String(s3.config.Bucket),
		Key:    aws.String(file),
	})

	urlStr, err := req.Presign(15 * time.Minute)

	return urlStr, nil
}

// MultiPart Upload
func (s3 s3) ObjectMultiPartUpload() {

}

// MultiPart Download
func (s3 s3) ObjectMultiPartDownload() {

}

// Bucket Objects List
func (s3 s3) BucketObjectList() {

	svc, err := s3.newSVCSession()

	// Get the list of items
	resp, err := svc.ListObjectsV2(&s3Pkg.ListObjectsV2Input{Bucket: aws.String("bucket")})
	if err != nil {

	}

	for _, item := range resp.Contents {
		fmt.Println("Name:         ", *item.Key)
		fmt.Println("Last modified:", *item.LastModified)
		fmt.Println("Size:         ", *item.Size)
		fmt.Println("Storage class:", *item.StorageClass)
		fmt.Println("")
	}

	fmt.Println("Found", len(resp.Contents), "items in bucket", "bucket")
	fmt.Println("")

}

// Object Metadata
func (s3 s3) ObjectMetadata() {

	svc, err := s3.newSVCSession()

	input := &s3Pkg.HeadObjectInput{
		Bucket: aws.String("bucket"),
		Key:    aws.String("obj"),
	}

	result, err := svc.HeadObject(input)
	if err != nil {
		// if aerr, ok := err.(awserr.Error); ok {
		// 	switch aerr.Code() {
		// 	default:
		// 		fmt.Println(aerr.Error())
		// 	}
		// } else {
		// 	// Print the error, cast err to awserr.Error to get the Code and
		// 	// Message from an error.
		// 	fmt.Println(err.Error())
		// }
		// return
	}

	fmt.Println(result)

}
