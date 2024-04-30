package aws

import (
    "context"

    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Service struct {
    DynamoDB *dynamodb.Client
}

func NewService() *Service {
    cfg, err := config.LoadDefaultConfig(context.Background())
    if err != nil {
        return nil
    }

    client := dynamodb.NewFromConfig(cfg)

    return &Service{
        DynamoDB: client,
    }
}
