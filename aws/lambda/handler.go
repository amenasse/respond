package main


import (
    "context"
    "github.com/amenasse/respond/statuscode"
    "github.com/aws/aws-lambda-go/events"
    "log"
    "os"
    "strconv"
    runtime "github.com/aws/aws-lambda-go/lambda"
)



func lambdaHandler(statusCode int) func(context.Context, events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	description := statuscode.Description[statusCode]
	if description == "" {
		description = "Unknown"
	}

        return func(ctx context.Context, r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

            host, header_present := r.Headers["Host"]
            if !header_present {
                host = "''"
            t
            log.Printf("%s %s %s %s", host, r.HTTPMethod, r.RequestContext.Protocol, r.Path)


            return events.APIGatewayProxyResponse{Body: description, StatusCode: statusCode}, nil
        }

}



func main() {
    status_code := 200
    if env_var := os.Getenv("RESPONSE_STATUS"); env_var != "" {
        var err error
        if status_code, err = strconv.Atoi(env_var); err != nil {
                log.Fatal("Illegal status code ")
        }
    }

    runtime.Start(lambdaHandler(status_code))
}
