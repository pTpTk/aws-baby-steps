
package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	fmt.Printf("Version: %s\n", request.Version)
	fmt.Print(request ,"\n")
	fmt.Printf("Processing request data for request %s.\n", request.RequestContext.RequestID)
	fmt.Printf("Body size = %d.\n", len(request.Body))

	fmt.Println("Headers:")
	for key, value := range request.Headers {
		fmt.Printf("    %s: %s\n", key, value)
	}

	return events.APIGatewayV2HTTPResponse{
		Body: "Hello world",
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "test/html",
		},
	}, nil
}

func main() {
	lambda.Start(handleRequest)
}
