package parser

import "github.com/valyala/fastjson"

type RefNode struct {
	Name string

	metadata *fastjson.Object
}

func IsRef(def *fastjson.Value) bool {
	return def != nil && def.Exists("ref")
}

func ParseRef(def *fastjson.Value) (RefNode, error) {
	name := string(def.GetStringBytes("ref"))
	metadata := def.GetObject("metadata")
	return RefNode{name, metadata}, nil
}

func (node RefNode) Metadata(key string) *fastjson.Value {
	if node.metadata == nil {
		return nil
	}

	return node.metadata.Get(key)
}
