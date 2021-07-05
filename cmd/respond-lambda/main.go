package main

import (
	"github.com/amenasse/respond/aws/lambda"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"log"
	"os"
	"strconv"
)

func main() {

	var code int = 200

	if env_var := os.Getenv("RESPONSE_STATUS"); env_var != "" {
		var err error
		if code, err = strconv.Atoi(env_var); err != nil {
			log.Fatal("Illegal status code ")
		}
	}

	if code < 200 || code > 599 {
		log.Fatal("Status code out of range (should be between 200-599)")
	}

	runtime.Start(lambda.ApiGatewayV2Handler(code))
}
