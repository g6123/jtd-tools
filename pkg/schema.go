package pkg

import (
	"github.com/g6123/jtd-tools/internal/parser"
	"github.com/g6123/jtd-tools/internal/schema"
)

func ToSchema(doc parser.DocumentNode) schema.RootSchema {
	return schema.FromDocument(doc)
}
