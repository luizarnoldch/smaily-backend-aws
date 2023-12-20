package e2e

import (
	"context"
	"log"
	"main/src/application/service"
	"main/src/domain/model"
	"main/src/infrastructure/repository_adapter"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"
)

func TestCreatePetFunctionally(t *testing.T) {
	// Using context.Background() as a base context is a common practice
	ctx := context.Background()

	// Initialize AWS configuration
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		t.Fatalf("Failed to load AWS configuration: %v", err)
	}

	// Create a DynamoDB client
	dynamoClient := dynamodb.NewFromConfig(cfg)

	// Initialize your repository
	repo := repository_adapter.NewDynamoPetRepository(dynamoClient, ctx, "Smaily-pets")

	// Initialize your service
	service := service.NewPetDynamoService(repo)

	// Define a test pet
	testPetRequest := &model.PetRequest{
		Name:          "Test Pet",
		Birth_Date:    "2021-01-01",
		Visit_Counter: 1,
	}

	// Perform the test: Create the pet
	createdPet, err := service.CreatePet(testPetRequest)
	assert.NoError(t, err, "Should create pet without error")
	assert.NotNil(t, createdPet, "Created pet should not be nil")

	log.Println(createdPet)

	// Additional assertions as necessary...

	// Clean up: delete the created pet from DynamoDB
}
