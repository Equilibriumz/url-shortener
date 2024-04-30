package repository

import (
    "github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var tableName = "shortened_url_table"

type Repository struct {
    client *dynamodb.Client
}

func NewRepository(client *dynamodb.Client) *Repository {
    return &Repository{
        client: client,
    }
}
