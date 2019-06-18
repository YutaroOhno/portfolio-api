// package storage

// import (
//     "github.com/aws/aws-sdk-go/aws"
//     "github.com/aws/aws-sdk-go/aws/session"
//     "github.com/aws/aws-sdk-go/aws/credentials"
//     "github.com/aws/aws-sdk-go/service/s3"
// )

// type Storage struct {
// 	Client *s3.S3
// }

// func NewStorage() *Storage {
// 	creds := credentials.NewStaticCredentials("AWS_ACCESS_KEY", "AWS_SECRET_ACCESS_KEY", "")
// 	sess, err := session.NewSession(&aws.Config{
// 		Credentials: creds,
// 		Region: aws.String("ap-northeast-1")},
// 	)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	svc := s3.New(sess)

// 	return  &Storage {
// 		Client: svc,
// 	}
// }

// func (storage *Storage) Upload() {

// }