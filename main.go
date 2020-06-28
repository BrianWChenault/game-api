package main

import (
	"errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ErrGameNotProvided = errors.New("no game id was provided")
)

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{}, nil
}

func main() {
	lambda.Start(Handler)
}
