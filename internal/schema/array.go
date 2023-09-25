package schema

import "github.com/g6123/jtd-tools/internal/parser"

type ArraySchema struct {
	Type  string `json:"type"`
	Items Schema `json:"items"`
}

func FromArray(node parser.ArrayNode) ArraySchema {
	return ArraySchema{Type: "array", Items: From(node.Element)}
}
