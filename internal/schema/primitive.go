package schema

import "github.com/g6123/jtd-tools/internal/parser"

type PrimitiveSchema struct {
	Type string `json:"type"`
}

type StringSchema struct {
	Type   string `json:"type"`
	Format string `json:"format,omitempty"`
}

type UnsignedIntegerSchema struct {
	Type    string `json:"type"`
	Minimum int    `json:"minimum"`
}

func FromPrimitive(node parser.PrimitiveNode) Schema {
	switch node.Type {
	case "int8":
	case "int16":
	case "int32":
	case "int64":
		return PrimitiveSchema{Type: "integer"}

	case "uint8":
	case "uint16":
	case "uint32":
	case "uint64":
		return UnsignedIntegerSchema{Type: "integer", Minimum: 0}

	case "float32":
	case "float64":
		return PrimitiveSchema{Type: "number"}

	case "timestamp":
		return StringSchema{Type: "string", Format: "date-time"}
	}

	return PrimitiveSchema{Type: node.Type}
}
