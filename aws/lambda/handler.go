package lambda

import (
	"context"
	"github.com/amenasse/respond/statuscode"
	"github.com/aws/aws-lambda-go/events"
	"log"
)

func ApiGatewayV1Handler(statusCode int) func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	description := statuscode.Description[statusCode]
	if description == "" {
		description = "Unknown"
	}

	return func(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

		host, header_present := r.Headers["Host"]
		if !header_present {
			host = "''"
		}
		log.Printf("%s %s %s %s", host, r.HTTPMethod, r.RequestContext.Protocol, r.Path)

		return events.APIGatewayProxyResponse{Body: description, StatusCode: statusCode}, nil
	}

}

func ApiGatewayV2Handler(statusCode int) func(context.Context, events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	description := statuscode.Description[statusCode]
	if description == "" {
		description = "Unknown"
	}

	return func(ctx context.Context, r events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

		host, header_present := r.Headers["host"]
		if !header_present {
			host = "''"
		}
		log.Printf("%s %s %s %s", host, r.RequestContext.HTTP.Method, r.RequestContext.HTTP.Protocol, r.RequestContext.HTTP.Path)

		return events.APIGatewayV2HTTPResponse{Body: description, StatusCode: statusCode}, nil
	}

}
