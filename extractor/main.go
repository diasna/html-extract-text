package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/PuerkitoBio/goquery"
)

func parse(url string, selector string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Print(err)
	}
	var result []string
	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		result = append(result, title)
	})
	return result, nil
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	if request.Headers["X-API-Key"] == os.Getenv("API_KEY") {
		url := request.QueryStringParameters["url"]
		selector := request.QueryStringParameters["selector"]
		var resp []string
		if len(url) < 1 && len(selector) < 1 {
			resp = []string{"insufficient parameters"}
		} else {
			resp, _ = parse(url, selector)
		}
		bytes, _ := json.Marshal(resp)
		return events.APIGatewayProxyResponse{Body: string(bytes), StatusCode: 200, Headers: map[string]string{"Content-Type": "application/json"}}, nil
	} else {
		return events.APIGatewayProxyResponse{StatusCode: 401}, nil
	}
}

func main() {
	lambda.Start(handler)
}
