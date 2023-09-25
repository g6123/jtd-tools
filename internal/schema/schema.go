package schema

import (
	"github.com/g6123/jtd-tools/internal/parser"
)

type Schema interface{}

type AnySchema struct{}

type AnyOfSchema struct {
	AnyOf []Schema `json:"anyOf"`
}

type ConstSchema struct {
	Const string `json:"const"`
}

func From(ctx SchemaContext, node parser.Node) Schema {
	switch typed_node := node.(type) {
	case parser.DocumentNode:
		return FromDocument(typed_node)

	case parser.ObjectNode:
		return FromObject(ctx, typed_node)

	case parser.DiscriminatedUnionNode:
		return FromDiscriminatedUnion(ctx, typed_node)

	case parser.RecordNode:
		return FromRecord(ctx, typed_node)

	case parser.ArrayNode:
		return FromArray(ctx, typed_node)

	case parser.EnumNode:
		return FromEnum(ctx, typed_node)

	case parser.RefNode:
		return FromRef(ctx, typed_node)

	case parser.PrimitiveNode:
		return FromPrimitive(ctx, typed_node)
	}

	return AnySchema{}
}

func NewNullableSchema(schema Schema) Schema {
	return AnyOfSchema{[]Schema{schema, BasicSchema{Type: "null"}}}
}
