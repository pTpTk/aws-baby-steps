package main

import (
	"errors"
	"fmt"
	"log"
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// DefaultHTTPGetAddress Default Address
	DefaultHTTPGetAddress = "https://checkip.amazonaws.com"

	// ErrNoIP No IP found in response
	ErrNoIP = errors.New("No IP in HTTP response")

	// ErrNon200Response non 200 status code in response
	ErrNon200Response = errors.New("Non 200 Response found")
)

type Message struct {
	V1 string `json:"key1"`
	V2 string `json:"key2"`
	V3 string `json:"key3"`
}

func handler(ctx context.Context, event json.RawMessage) (events.APIGatewayProxyResponse, error) {
	var message Message
	if err := json.Unmarshal(event, &message); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return events.APIGatewayProxyResponse{}, err
	}


	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello, %s", message.V1),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}