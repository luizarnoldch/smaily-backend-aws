package pets_test

import (
	"context"
	"log"
	"main/src/infrastructure"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/suite"
)

var (
	TABLE_NAME = os.Getenv("TABLE_NAME")
)

type DynamoPetRepositorySuite struct {
    suite.Suite
    infrastructure.PetRepository
}

func (suite *DynamoPetRepositorySuite) SetupTest() {
    var ctx context.Context
    cfg, err := config.LoadDefaultConfig(ctx)
    suite.NoError(err)
	dynamoClient := dynamodb.NewFromConfig(cfg)
    dynamoPetRepository := infrastructure.NewDynamoPetRepository(dynamoClient, ctx, TABLE_NAME)

    suite.PetRepository = dynamoPetRepository
}

func (suite *DynamoPetRepositorySuite) TestGetAllPets() {
    pets, err := suite.GetAllPets()
    suite.NoError(err)
    for _, pet := range pets {
        log.Println(pet.ID)
    }
}

func TestDynamoPetRepositorySuite (t *testing.T) {
    suite.Run(t, new(DynamoPetRepositorySuite))
}