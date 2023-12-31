package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"main/src/application/service"
	"main/src/domain/model"
	"main/src/infrastructure/repository_adapter"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	TABLE_NAME = os.Getenv("TABLE_NAME")
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	log.Println("create lamba starts")

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf("%s", err), StatusCode: 502}, err
	}

	dynamoClient := dynamodb.NewFromConfig(cfg)

	dynamoInfrastructure := repository_adapter.NewDynamoPetRepository(dynamoClient, ctx, TABLE_NAME)
	dynamoService := service.NewPetDynamoService(dynamoInfrastructure)

	var petRequest model.PetRequest

	log.Println("unmarshal the content body")
	if err := json.Unmarshal([]byte(request.Body), &petRequest); err != nil {
		log.Println("error parsing request body as json:", err)
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf("error parsing request body as json: %s", err), StatusCode: 502}, nil
	}

	log.Println("getting all pets")
	response, err := dynamoService.GetAllPets()
	if err != nil {
		log.Println("Error getting all the pet:", err)
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf("Error getting all the pet: %s", err), StatusCode: 502}, nil
	}

	log.Println("converting the response to json")
	responseBody, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error marshaling response to JSON: %s", err)
		return events.APIGatewayProxyResponse{Body: fmt.Sprintf("Error marshaling response to JSON: %s", err), StatusCode: 502}, nil
	}

	headers := map[string]string{
		"Access-Control-Allow-Origin":  "*",
		"Access-Control-Allow-Methods": "POST,OPTIONS,DELETE,GET,HEAD,PUT",
		"Access-Control-Allow-Headers": "Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Custom-Header",
		"Content-Type":                 "application/json",
	}

	log.Println("pet was created")
	return events.APIGatewayProxyResponse{
		Headers:    headers,
		Body:       string(responseBody),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handler)
}
