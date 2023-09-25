package parser

import "github.com/valyala/fastjson"

type PrimitiveNode struct {
	Type     string
	Nullable bool

	metadata *fastjson.Object
}

func IsPrimitive(def *fastjson.Value) bool {
	return def != nil && def.Exists("type")
}

func ParsePrimitive(def *fastjson.Value) (PrimitiveNode, error) {
	primitive_type := string(def.GetStringBytes("type"))
	nullable := def.GetBool("nullable")
	metadata := def.GetObject("metadata")
	return PrimitiveNode{primitive_type, nullable, metadata}, nil
}

func (node PrimitiveNode) Metadata(key string) *fastjson.Value {
	if node.metadata == nil {
		return nil
	}

	return node.metadata.Get(key)
}
