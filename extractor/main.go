package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       os.Getenv("API_KEY"),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
