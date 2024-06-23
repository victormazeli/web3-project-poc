package aws

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func AwsSetup() {
    // Specify your AWS credentials and region
    awsRegion := "your-aws-region"
    awsAccessKeyID := "your-access-key-id"
    awsSecretAccessKey := "your-secret-access-key"

    // Create a new AWS session
    sess, err := session.NewSession(&aws.Config{
        Region:      aws.String(awsRegion),
        Credentials: credentials.NewStaticCredentials(awsAccessKeyID, awsSecretAccessKey, ""),
    })
    if err != nil {
        fmt.Println("Error creating session:", err)
        return
    }

    // Specify the bucket name and file path
    bucketName := "your-bucket-name"
    filePath := "/path/to/your/file.txt"

    // Open the file to upload
    file, err := os.Open(filePath)
    if err != nil {
        fmt.Println("Error opening file:", err)
        return
    }
    defer file.Close()

    // Create an S3 client
    svc := s3.New(sess)

    // Upload file to S3
    _, err = svc.PutObject(&s3.PutObjectInput{
        Bucket: aws.String(bucketName),
        Key:    aws.String(filepath.Base(filePath)),
        Body:   file,
    })
    if err != nil {
        fmt.Println("Error uploading file:", err)
        return
    }

    fmt.Println("File uploaded successfully.")
}
