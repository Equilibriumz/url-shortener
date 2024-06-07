package util

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func AttributeValueToData(value types.AttributeValue) any {
	switch v := value.(type) {
	case *types.AttributeValueMemberB:
		return v.Value // Value is []byte

	case *types.AttributeValueMemberBOOL:
		return v.Value // Value is bool

	case *types.AttributeValueMemberBS:
		return v.Value // Value is [][]byte

	case *types.AttributeValueMemberL:
		return v.Value // Value is []types.AttributeValue

	case *types.AttributeValueMemberM:
		return v.Value // Value is map[string]types.AttributeValue

	case *types.AttributeValueMemberN:
		return v.Value // Value is string

	case *types.AttributeValueMemberNS:
		return v.Value // Value is []string

	case *types.AttributeValueMemberNULL:
		return v.Value // Value is bool

	case *types.AttributeValueMemberS:
		return v.Value // Value is string

	case *types.AttributeValueMemberSS:
		return v.Value // Value is []string

	case *types.UnknownUnionMember:
		return v.Value // Value type []byte

	default:
		return nil
	}
}
