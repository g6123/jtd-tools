package parser

import "github.com/valyala/fastjson"

type ArrayNode struct {
	Element  Node
	Nullable bool

	metadata *fastjson.Object
}

func IsArray(def *fastjson.Value) bool {
	return def != nil && def.Exists("elements")
}

func ParseArray(def *fastjson.Value) (ArrayNode, error) {
	element, err := Parse(def.Get("elements"))
	nullable := def.GetBool("nullable")
	metadata := def.GetObject("metadata")
	return ArrayNode{element, nullable, metadata}, err
}

func (node ArrayNode) Metadata(key string) *fastjson.Value {
	if node.metadata == nil {
		return nil
	}

	return node.metadata.Get(key)
}
