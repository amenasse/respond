package lambda

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/amenasse/respond/cmd"
        "net/http"
)


func ApiGatewayV1Handler(statusCode int) func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	body := cmd.Body(statusCode)
	return func(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

                headers := make(http.Header)
                for key, value := range r.Headers {
                        headers.Add(key, value)
                }
		cmd.Log(headers, r.HTTPMethod, r.RequestContext.Protocol, r.Path)
		return events.APIGatewayProxyResponse{Body: body, StatusCode: statusCode}, nil
	}

}

func ApiGatewayV2Handler(statusCode int) func(context.Context, events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	body := cmd.Body(statusCode)
	return func(ctx context.Context, r events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

                headers := make(http.Header)
                for key, value := range r.Headers {
                        headers.Add(key, value)
                }
		cmd.Log(headers, r.RequestContext.HTTP.Method, r.RequestContext.HTTP.Protocol, r.RequestContext.HTTP.Path)
		return events.APIGatewayV2HTTPResponse{Body: body, StatusCode: statusCode}, nil
	}

}
