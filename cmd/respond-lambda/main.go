package main

import (
	"github.com/amenasse/respond/aws/lambda"
	"github.com/amenasse/respond/cmd"
	runtime "github.com/aws/aws-lambda-go/lambda"
)

func main() {

	code := cmd.GetStatusCode()
	runtime.Start(lambda.ApiGatewayV2Handler(code))
}
