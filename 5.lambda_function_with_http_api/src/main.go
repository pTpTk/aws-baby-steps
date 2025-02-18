package main

import (
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"
)

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Print("Resource:              ", req.Resource, "\n")
	fmt.Print("Path:                  ", req.Path, "\n")
	fmt.Print("HTTPMethod:            ", req.HTTPMethod, "\n")
	fmt.Print("Headers:               ", req.Headers, "\n")
	fmt.Print("QueryStringParameters: ", req.QueryStringParameters, "\n")
	fmt.Print("PathParameters:        ", req.PathParameters, "\n")
	fmt.Print("StageVariables:        ", req.StageVariables, "\n")
	fmt.Print("RequestContext:        ", req.RequestContext, "\n")
	fmt.Print("Body:                  ", req.Body, "\n")
	fmt.Print("IsBase64Encoded:       ", req.IsBase64Encoded, "\n")

	return events.APIGatewayProxyResponse{}, nil
}

func main() {
	lambda.Start(handler)
}