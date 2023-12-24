package adapter

import (
	"balance-inquiry/internal/domain"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type DynamoDBRepository struct {
	DynamoDBClient DynamoDBAPI
	TableName      string
}

type DynamoDBAPI interface {
	Scan(*dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
}

func NewDynamoDBRepository(tableName string, client DynamoDBAPI) *DynamoDBRepository {
	return &DynamoDBRepository{
		DynamoDBClient: client,
		TableName:      tableName,
	}
}

func (r *DynamoDBRepository) GetPointByUser(id string) (*[]domain.Point, error) {
	filt := expression.Name("user").Contains(id)
	expr, err := expression.NewBuilder().
		WithFilter(filt).
		Build()

	if err != nil {
		fmt.Println(err)
	}
	input := &dynamodb.ScanInput{
		TableName:        aws.String(r.TableName),
		FilterExpression: expr.Filter(),
	}

	result, err := r.DynamoDBClient.Scan(input)
	if err != nil {
		return nil, err
	}

	response := []domain.Point{}
	for _, item := range result.Items {
		var point domain.Point
		if err := dynamodbattribute.UnmarshalMap(item, &point); err != nil {
			return nil, err
		}

		response = append(response, point)
	}

	return &response, nil
}
