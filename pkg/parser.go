package pkg

import (
	"os"

	"github.com/g6123/jtd-tools/internal/parser"
	"github.com/g6123/jtd-tools/internal/schema"
	"github.com/valyala/fastjson"
)

func ParseFile(filename string) (*parser.DocumentNode, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return ParseBytes(file)
}

func ParseBytes(bytes []byte) (*parser.DocumentNode, error) {
	value, err := fastjson.ParseBytes(bytes)
	if err != nil {
		return nil, err
	}

	doc, err := parser.ParseDocument(value)
	return &doc, err
}

func ToSchema(doc parser.DocumentNode) schema.RootSchema {
	return schema.FromDocument(doc)
}
