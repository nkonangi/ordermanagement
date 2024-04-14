package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("hello world handler ")

	res := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}
	return res, nil
}
