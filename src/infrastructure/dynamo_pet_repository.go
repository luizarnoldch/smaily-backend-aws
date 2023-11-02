package infrastructure

import (
	"context"
	"fmt"
	"main/src/domain"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBPetRepository struct {
	client *dynamodb.Client
	ctx    context.Context
	table  string
}

func NewDynamoPetRepository(config aws.Config, ctx context.Context, table string) *DynamoDBPetRepository {
	return &DynamoDBPetRepository{
		client: dynamodb.NewFromConfig(config),
		ctx:    ctx,
		table:  table,
	}
}

func (repo *DynamoDBPetRepository) GetAllPets() ([]domain.PetResponse, error) {
	input := &dynamodb.ExecuteStatementInput{
		Statement: aws.String(fmt.Sprintf("SELECT * FROM \"%v\"", repo.table)),
	}

	result, err := repo.client.ExecuteStatement(context.TODO(), input)
	if err != nil {
		return []domain.PetResponse{}, err
	}

	var pets []domain.Pet
	err = attributevalue.UnmarshalListOfMaps(result.Items, &pets)
	if err != nil {
		return []domain.PetResponse{}, err
	}

	var response []domain.PetResponse
	for _, pet := range pets {
		petResponse := pet.ToPetResponse()
		response = append(response, petResponse)
	}

	return response, nil
}

func (repo *DynamoDBPetRepository) CreatePet(req *domain.Pet) (*domain.PetResponse, error) {
	item, err := attributevalue.MarshalMap(&req)
	if err != nil {
		return &domain.PetResponse{Message: err.Error()}, err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(repo.table),
		Item:  item,
	}

	output, err := repo.client.PutItem(repo.ctx, input)
	if err != nil {
		return &domain.PetResponse{Message: err.Error()}, err
	}

	var petResponse *domain.PetResponse
	err = attributevalue.UnmarshalMap(output.Attributes, &petResponse)
	if err != nil {
		return &domain.PetResponse{Message: err.Error()}, err
	}

	return petResponse, nil
}


