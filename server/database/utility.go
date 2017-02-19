package main

import (
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func getdbSession(devenv bool) (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession()
	if err != nil {
		return nil, err
	}
	if devenv {
		return dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:8400"), Region: aws.String("us-east-1"), Credentials: credentials.NewStaticCredentials("123", "123", "")}), nil
	}
	return dynamodb.New(sess), nil
}

func setStatus(w http.ResponseWriter, statusCode int, statusText string) {
	w.WriteHeader(statusCode)

}
