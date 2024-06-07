package repository

import (
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type Record struct {
	HashID    string `dynamodbav:"hash_id"`
	URLString string `dynamodbav:"url"`
}

// DTO

func DynamoDBObjectToRecord(item map[string]types.AttributeValue) *Record {
	if item == nil {
		return nil
	}

	record := new(Record)

	if err := attributevalue.UnmarshalMap(item, &record); err != nil {
		return nil
	}

	return record
}

func RecordToDynamoDBObject(record *Record) map[string]types.AttributeValue {
	if record == nil {
		return nil
	}

	item, err := attributevalue.MarshalMap(record)
	if err != nil {
		return nil
	}

	return item
}
