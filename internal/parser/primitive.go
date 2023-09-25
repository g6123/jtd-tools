package parser

import "github.com/valyala/fastjson"

type PrimitiveNode struct {
	Type string

	metadata *fastjson.Object
}

func IsPrimitive(def *fastjson.Value) bool {
	return def != nil && def.Exists("type")
}

func ParsePrimitive(def *fastjson.Value) (PrimitiveNode, error) {
	return PrimitiveNode{Type: string(def.GetStringBytes("type"))}, nil
}

func (node PrimitiveNode) Metadata(key string) *fastjson.Value {
	if node.metadata == nil {
		return nil
	}

	return node.metadata.Get(key)
}
