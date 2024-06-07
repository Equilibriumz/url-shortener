package repository

import (
	"reflect"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TestDynamoDBObjectToRecord(t *testing.T) {
	type args struct {
		item map[string]types.AttributeValue
	}
	tests := []struct {
		name string
		args args
		want *Record
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DynamoDBObjectToRecord(tt.args.item); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DynamoDBObjectToRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewRepository(t *testing.T) {
	type args struct {
		client *dynamodb.Client
	}
	tests := []struct {
		name string
		args args
		want *Repository
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRecordToDynamoDBObject(t *testing.T) {
	type args struct {
		record *Record
	}
	tests := []struct {
		name string
		args args
		want map[string]types.AttributeValue
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RecordToDynamoDBObject(tt.args.record); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RecordToDynamoDBObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetByKey(t *testing.T) {
	type fields struct {
		client *dynamodb.Client
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Record
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				client: tt.fields.client,
			}
			got, err := r.GetByKey(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetByKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByKey() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_InsertItem(t *testing.T) {
	type fields struct {
		client *dynamodb.Client
	}
	type args struct {
		record *Record
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				client: tt.fields.client,
			}
			if err := r.InsertItem(tt.args.record); (err != nil) != tt.wantErr {
				t.Errorf("InsertItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
