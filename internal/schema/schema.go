package schema

import (
	"github.com/g6123/jtd-tools/internal/parser"
)

type Schema interface{}

type AnySchema struct{}

func From(node parser.Node) Schema {
	switch typed_node := node.(type) {
	case parser.DocumentNode:
		return FromDocument(typed_node)

	case parser.ObjectNode:
		return FromObject(typed_node)

	case parser.DiscriminatedUnionNode:
		return FromDiscriminatedUnion(typed_node)

	case parser.RecordNode:
		return FromRecord(typed_node)

	case parser.ArrayNode:
		return FromArray(typed_node)

	case parser.EnumNode:
		return FromEnum(typed_node)

	case parser.RefNode:
		return FromRef(typed_node)

	case parser.PrimitiveNode:
		return FromPrimitive(typed_node)
	}

	return AnySchema{}
}
