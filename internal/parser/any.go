package parser

import "github.com/valyala/fastjson"

type AnyNode struct {
	metadata *fastjson.Object
}

func IsAny(def *fastjson.Value) bool {
	return def != nil && def.Type() == fastjson.TypeObject
}

func ParseAny(def *fastjson.Value) (AnyNode, error) {
	return AnyNode{def.GetObject("metadata")}, nil
}

func (node AnyNode) Metadata(key string) *fastjson.Value {
	if node.metadata == nil {
		return nil
	}

	return node.metadata.Get(key)
}
