package schema

import (
	"github.com/g6123/jtd-tools/internal/parser"
)

type BasicSchema struct {
	Type string `json:"type"`
}

type AnyOfBasicSchema struct {
	Type []string `json:"type"`
}

type StringSchema struct {
	Type   string `json:"type"`
	Format string `json:"format,omitempty"`
}

type UnsignedIntegerSchema struct {
	Type    string `json:"type"`
	Minimum int    `json:"minimum"`
}

func FromPrimitive(ctx SchemaContext, node parser.PrimitiveNode) Schema {
	var schema Schema

	switch node.Type {
	case "int8", "int16", "int32", "int64":
		schema = BasicSchema{Type: "integer"}

	case "uint8", "uint16", "uint32", "uint64":
		schema = UnsignedIntegerSchema{Type: "integer", Minimum: 0}

	case "float32", "float64":
		schema = BasicSchema{Type: "number"}

	case "timestamp":
		schema = StringSchema{Type: "string", Format: "date-time"}

	default:
		schema = BasicSchema{Type: node.Type}
	}

	if node.Nullable {
		switch s := schema.(type) {
		case BasicSchema:
			schema = AnyOfBasicSchema{[]string{s.Type, "null"}}

		default:
			schema = NewNullableSchema(s)
		}
	}

	return schema
}
