package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/amenasse/respond/cmd"
)



func ApiGatewayV1Handler(statusCode int) func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	body := cmd.Body(statusCode)
	return func(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		cmd.Log(r.Headers, r.HTTPMethod, r.RequestContext.Protocol, r.Path)
		return events.APIGatewayProxyResponse{Body: body, StatusCode: statusCode}, nil
	}

}

func ApiGatewayV2Handler(statusCode int) func(context.Context, events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	body := cmd.Body(statusCode)
	return func(ctx context.Context, r events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
		cmd.Log(r.Headers, r.RequestContext.HTTP.Method, r.RequestContext.HTTP.Protocol, r.RequestContext.HTTP.Path)
		return events.APIGatewayV2HTTPResponse{Body: body, StatusCode: statusCode}, nil
	}

}
