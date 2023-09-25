package schema

import (
	"github.com/g6123/jtd-tools/internal/parser"
)

type DiscriminatedUnionSchema struct {
	Type  string         `json:"type"`
	OneOf []ObjectSchema `json:"oneOf"`
}

func FromDiscriminatedUnion(ctx SchemaContext, node parser.DiscriminatedUnionNode) Schema {
	schema := DiscriminatedUnionSchema{
		Type:  "object",
		OneOf: make([]ObjectSchema, 0, len(node.Mapping)),
	}

	for variant_name, variant := range node.Mapping {
		variant_schema := FromObjectStrict(ctx, variant.Value)
		variant_schema.Properties[node.Disciriminator] = ConstSchema{variant_name}
		variant_schema.Required = append(variant_schema.Required, node.Disciriminator)
		schema.OneOf = append(schema.OneOf, variant_schema)
	}

	if node.Nullable {
		return NewNullableSchema(schema)
	} else {
		return schema
	}
}
