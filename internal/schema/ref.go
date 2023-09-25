package schema

import (
	"fmt"

	"github.com/g6123/jtd-tools/internal/parser"
)

type RefSchema struct {
	Ref string `json:"$ref"`
}

func FromRef(node parser.RefNode) RefSchema {
	return RefSchema{fmt.Sprintf("#/$defs/%s", node.Name)}
}
