package schema

import "github.com/g6123/jtd-tools/internal/parser"

type DiscriminatedUnionSchema struct {
	Type  string         `json:"type"`
	OneOf []ObjectSchema `json:"oneOf"`
}

type ConstSchema struct {
	Const string `json:"const"`
}

func FromDiscriminatedUnion(node parser.DiscriminatedUnionNode) DiscriminatedUnionSchema {
	schema := DiscriminatedUnionSchema{
		Type:  "object",
		OneOf: make([]ObjectSchema, len(node.Mapping)),
	}

	for variant_name, variant := range node.Mapping {
		variant_schema := FromObject(variant.Value)
		variant_schema.Properties[node.Disciriminator] = ConstSchema{variant_name}
		variant_schema.Required = append(variant_schema.Required, node.Disciriminator)
		schema.OneOf = append(schema.OneOf, variant_schema)
	}

	return schema
}
