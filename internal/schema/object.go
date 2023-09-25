package schema

import "github.com/g6123/jtd-tools/internal/parser"

type ObjectSchema struct {
	Type       string            `json:"type"`
	Properties map[string]Schema `json:"properties"`
	Required   []string          `json:"required"`
}

func FromObject(node parser.ObjectNode) ObjectSchema {
	schema := ObjectSchema{
		Type:       "object",
		Properties: make(map[string]Schema, len(node.Properties)),
	}

	for _, prop := range node.Properties {
		schema.Properties[prop.Key] = From(prop.Value)

		if !prop.IsOptional {
			schema.Required = append(schema.Required, prop.Key)
		}
	}

	return schema
}
