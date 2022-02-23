package main

import (
	"github.com/amenasse/respond/cmd"
	runtime "github.com/aws/aws-lambda-go/lambda"
        "log"
        "os"
)

func main() {

	code := cmd.GetStatusCode()
        handler_name := "api-v2"
        if env_var := os.Getenv("RESPONSE_LAMBDA_HANDLER"); env_var != "" {
            handler_name = env_var
        }

	if handler_name == "api-v2" {
		runtime.Start(ApiGatewayV2Handler(code))
	} else if handler_name == "api-v1" {

		runtime.Start(ApiGatewayV1Handler(code))
	} else
        {
            log.Fatalf("Unknown handler: %s",handler_name)
    }
}
