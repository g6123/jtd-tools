package parser

import "github.com/valyala/fastjson"

type RefNode struct {
	Name     string
	Nullable bool

	metadata *fastjson.Object
}

func IsRef(def *fastjson.Value) bool {
	return def != nil && def.Exists("ref")
}

func ParseRef(def *fastjson.Value) (RefNode, error) {
	name := string(def.GetStringBytes("ref"))
	nullable := def.GetBool("nullable")
	metadata := def.GetObject("metadata")
	return RefNode{name, nullable, metadata}, nil
}

func (node RefNode) Metadata(key string) *fastjson.Value {
	if node.metadata == nil {
		return nil
	}

	return node.metadata.Get(key)
}
