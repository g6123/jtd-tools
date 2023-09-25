package schema

import (
	"github.com/g6123/jtd-tools/internal/parser"
)

type ArraySchema struct {
	Type  string `json:"type"`
	Items Schema `json:"items"`
}

func FromArray(ctx SchemaContext, node parser.ArrayNode) Schema {
	schema := ArraySchema{
		Type:  "array",
		Items: From(ctx, node.Element),
	}

	if node.Nullable {
		return NewNullableSchema(schema)
	} else {
		return schema
	}
}
