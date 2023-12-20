package repository_adapter

import (
	"context"
	"fmt"
	"main/src/domain/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBPetRepository struct {
	client *dynamodb.Client
	ctx    context.Context
	table  string
}

func NewDynamoPetRepository(client *dynamodb.Client, ctx context.Context, table string) *DynamoDBPetRepository {
	return &DynamoDBPetRepository{
		client: client,
		ctx:    ctx,
		table:  table,
	}
}

func (repo *DynamoDBPetRepository) GetAllPets() ([]model.PetResponse, error) {
	input := &dynamodb.ExecuteStatementInput{
		Statement: aws.String(fmt.Sprintf("SELECT * FROM \"%v\"", repo.table)),
	}

	result, err := repo.client.ExecuteStatement(context.TODO(), input)
	if err != nil {
		return []model.PetResponse{}, err
	}

	var response []model.PetResponse
	err = attributevalue.UnmarshalListOfMaps(result.Items, &response)
	if err != nil {
		return []model.PetResponse{}, err
	}

	return response, nil
}

func (repo *DynamoDBPetRepository) CreatePet(req *model.Pet) (*model.PetResponse, error) {
	item, err := attributevalue.MarshalMap(&req)
	if err != nil {
		return &model.PetResponse{Message: err.Error()}, err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(repo.table),
		Item:      item,
	}

	output, err := repo.client.PutItem(repo.ctx, input)
	if err != nil {
		return &model.PetResponse{Message: err.Error()}, err
	}

	var petResponse *model.PetResponse
	err = attributevalue.UnmarshalMap(output.Attributes, &petResponse)
	if err != nil {
		return &model.PetResponse{Message: err.Error()}, err
	}

	return petResponse, nil
}

func (repo *DynamoDBPetRepository) DeletePet(petId string) (*model.PetResponse, error) {
	key, err := attributevalue.MarshalMap(map[string]string{"PetId": petId})
	if err != nil {
		return &model.PetResponse{Message: err.Error()}, err
	}

	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(repo.table),
		Key:       key,
	}

	result, err := repo.client.DeleteItem(repo.ctx, input)
	if err != nil {
		return &model.PetResponse{Message: err.Error()}, err
	}

	var petResponse model.PetResponse
	err = attributevalue.UnmarshalMap(result.Attributes, &petResponse)
	if err != nil {
		return &model.PetResponse{Message: err.Error()}, err
	}

	return &petResponse, nil
}
