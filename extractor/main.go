package main

import (
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.Headers["X-API-Key"] == os.Getenv("API_KEY") {
		return events.APIGatewayProxyResponse{
			Body:       os.Getenv("API_KEY"),
			StatusCode: 200,
		}, nil
	} else {
		return events.APIGatewayProxyResponse{
			StatusCode: 401,
		}, nil
	}
}

func main() {
	lambda.Start(handler)
}
