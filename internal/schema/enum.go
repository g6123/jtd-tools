package schema

import "github.com/g6123/jtd-tools/internal/parser"

type EnumSchema struct {
	Enum []string `json:"enum"`
}

func FromEnum(node parser.EnumNode) EnumSchema {
	return EnumSchema{node.Cases}
}
