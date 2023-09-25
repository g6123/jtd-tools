package schema

import (
	"encoding/json"

	"github.com/g6123/jtd-tools/internal/parser"
)

type RootSchema struct {
	Id string
	Schema
	defs map[string]Schema
}

func FromDocument(node parser.DocumentNode) RootSchema {
	schema_id_bytes, _ := node.Metadata("schemaId").StringBytes()
	schema_id := string(schema_id_bytes)

	ctx := SchemaContext{
		SchemaId: schema_id,
	}

	schema := RootSchema{
		Schema: From(ctx, node.Root),
		Id:     string(schema_id),
		defs:   make(map[string]Schema, len(node.Definitions)),
	}

	for def_name, def_node := range node.Definitions {
		schema.defs[def_name] = From(ctx, def_node.Value)
	}

	return schema
}

func (schema RootSchema) MarshalJSON() ([]byte, error) {
	root := make(map[string]interface{})

	if schema.Id != "" {
		root["$id"] = schema.Id
	}

	root["$defs"] = schema.defs

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
