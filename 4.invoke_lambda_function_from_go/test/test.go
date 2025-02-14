package main

import (
	"context"
	"fmt"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
)

type Req struct {
	V1 string `json:"key1"`
	V2 string `json:"key2"`
	V3 string `json:"key3"`
}

// main uses the AWS SDK for Go (v2) to create an AWS Lambda client and list up to 10
// functions in your account.
// This example uses the default settings specified in your shared credentials
// and config files.
func main() {
	ctx := context.Background()
	sdkConfig, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
		return
	}
	lambdaClient := lambda.NewFromConfig(sdkConfig)

 	req := Req{
 		V1: "V1",
 		V2: "value2",
   		V3: "value3",
 	}

	payload, err := json.Marshal(req)

	output, err := lambdaClient.Invoke(ctx, &lambda.InvokeInput{
		FunctionName: aws.String("myFunction"),
		LogType:      types.LogTypeNone,
		Payload:      payload,
	})
	fmt.Printf("%s", output.Payload)
}
