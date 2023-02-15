package main

/*
build and zip for lambda aws
export GOARCH=amd64
export GOOS=linux
go build -ldflags="-s -w" main.go
zip  list3.zip main
*/
import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type MyEvent struct {
	Name string `json:"name"`
}

func getListBucket(client *s3.S3) (*s3.ListBucketsOutput, error) {
	res, err := client.ListBuckets(nil)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func HandleRequest() (string, error) {

	s3Client := s3.New(session.New())
	listBucket, err := getListBucket(s3Client)
	if err != nil {
		return fmt.Sprintf("Error get list buckets with err: %v", err), err
	}
	strOut := ""
	for _, bucket := range listBucket.Buckets {
		strOut = strOut + fmt.Sprintf("Found bucket: %s, created at: %s\n", *bucket.Name, *bucket.CreationDate)
	}
	return strOut, err
}

func main() {
	lambda.Start(HandleRequest)
}
