package s3

import (
	// "fmt"

	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/awserr"
	s3Pkg "github.com/aws/aws-sdk-go/service/s3"

)

// Create Bucket

//Bucket List

// Check Bucket is Exist

// Bucket CORS List
func (s3 s3) IndexBucketCORS(bucket string) (any, error) {

	svc, err := s3.newSVCSession()

	input := &s3Pkg.GetBucketCorsInput{
		Bucket: aws.String(bucket),
	}

	result, err := svc.GetBucketCors(input)
	if err != nil {

		return nil, err
		// if aerr, ok := err.(awserr.Error); ok {
		// 	switch aerr.Code() {
		// 	default:
		// 		fmt.Println(aerr.Error())
		// 	}
		// } else {
		// 	fmt.Println(err.Error())
		// }
		// return
	}

	return result, err

}

//Create cors

//Delete cors

// Delete Tag
func (s3 s3) DeleteBucketTag() error {
	return nil
}

//Add Tag
