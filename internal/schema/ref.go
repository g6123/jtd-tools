package schema

import (
	"fmt"

	"github.com/g6123/jtd-tools/internal/parser"
)

type RefSchema struct {
	Ref string `json:"$ref"`
}

func FromRef(ctx SchemaContext, node parser.RefNode) Schema {
	schema := RefSchema{fmt.Sprintf("%s#/$defs/%s", ctx.SchemaId, node.Name)}

	if node.Nullable {
		return NewNullableSchema(schema)
	} else {
		return schema
	}
}
