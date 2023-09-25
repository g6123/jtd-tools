package schema

import (
	"github.com/g6123/jtd-tools/internal/parser"
)

type EnumSchema struct {
	Enum []string `json:"enum"`
}

func FromEnum(ctx SchemaContext, node parser.EnumNode) Schema {
	schema := EnumSchema{node.Cases}

	if node.Nullable {
		return NewNullableSchema(schema)
	} else {
		return schema
	}
}
