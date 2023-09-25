package schema

import "github.com/g6123/jtd-tools/internal/parser"

type RecordSchema struct {
	Type                 string `json:"type"`
	AdditionalProperties Schema `json:"additionalProperties"`
}

func FromRecord(ctx SchemaContext, node parser.RecordNode) Schema {
	schema := RecordSchema{
		Type:                 "object",
		AdditionalProperties: From(ctx, node.Value),
	}

	if node.Nullable {
		return NewNullableSchema(schema)
	} else {
		return schema
	}
}
