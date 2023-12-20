package test

import (
	"context"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestDynamoDBConnection(t *testing.T) {
	// Cargar la configuración del SDK de AWS
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithEndpointResolverWithOptions(
			aws.EndpointResolverWithOptionsFunc(
				func(service, region string, options ...interface{}) (aws.Endpoint, error) {
					return aws.Endpoint{URL: "http://localhost:8000"}, nil
				},
			),
		),
	)
	if err != nil {
		t.Fatalf("Falló la configuración: %s", err)
	}

	// Crear un cliente DynamoDB con la configuración
	svc := dynamodb.NewFromConfig(cfg)

	// Nombre de la tabla para la prueba
	tableName := "TestTable"

	// Crear una tabla
	_, err = svc.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("ID"),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("ID"),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   aws.String(tableName),
		BillingMode: types.BillingModePayPerRequest,
	})
	if err != nil {
		t.Fatalf("No se pudo crear la tabla: %s", err)
	}

	// Esperar a que la tabla sea creada
	time.Sleep(5 * time.Second)

	// Llamar a la operación ListTables
	resp, err := svc.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		t.Fatalf("No se pudo conectar a DynamoDB: %s", err)
	}

	// Comprobar si la tabla creada está en la lista
	found := false
	for _, tName := range resp.TableNames {
		if tName == tableName {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("La tabla creada '%s' no se encontró en la lista", tableName)
	}
}
