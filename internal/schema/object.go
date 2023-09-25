package schema

import (
	"github.com/g6123/jtd-tools/internal/parser"
)

type ObjectSchema struct {
	Type       string            `json:"type"`
	Properties map[string]Schema `json:"properties"`
	Required   []string          `json:"required"`
}

func FromObject(ctx SchemaContext, node parser.ObjectNode) Schema {
	schema := FromObjectStrict(ctx, node)

	if node.Nullable {
		return NewNullableSchema(schema)
	} else {
		return schema
	}
}

func FromObjectStrict(ctx SchemaContext, node parser.ObjectNode) ObjectSchema {
	schema := ObjectSchema{
		Type:       "object",
		Properties: make(map[string]Schema, len(node.Properties)),
	}

	for _, prop := range node.Properties {
		schema.Properties[prop.Key] = From(ctx, prop.Value)

		if !prop.IsOptional {
			schema.Required = append(schema.Required, prop.Key)
		}
	}

	return schema
}
