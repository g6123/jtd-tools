package schema

import (
	"encoding/json"

	"github.com/g6123/jtd-tools/internal/parser"
)

type RootSchema struct {
	Schema
	Defs map[string]Schema
}

func FromDocument(node parser.DocumentNode) RootSchema {
	schema := RootSchema{
		Schema: From(node.Root),
		Defs:   make(map[string]Schema, len(node.Definitions)),
	}

	for def_name, def_node := range node.Definitions {
		schema.Defs[def_name] = From(def_node.Value)
	}

	return schema
}

func (schema RootSchema) MarshalJSON() ([]byte, error) {
	root := make(map[string]interface{})
	root["$defs"] = schema.Defs

	base, err := json.Marshal(schema.Schema)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(base, &root)
	if err != nil {
		return nil, err
	}

	return json.Marshal(root)
}
