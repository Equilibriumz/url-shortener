package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func (r *Repository) GetByKey(key string) (*Record, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"hash_id": &types.AttributeValueMemberS{
				Value: key,
			},
		},
		TableName: &tableName,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	output, err := r.client.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}

	fmt.Println(output.Item)

	return DynamoDBObjectToRecord(output.Item), nil
}

func (r *Repository) InsertItem(record *Record) error {
	input := &dynamodb.PutItemInput{
		Item:      RecordToDynamoDBObject(record),
		TableName: &tableName,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	_, err := r.client.PutItem(ctx, input)

	return err
}
