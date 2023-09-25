package parser

import "github.com/valyala/fastjson"

type RecordNode struct {
	Value Node

	metadata *fastjson.Object
}

func IsRecord(def *fastjson.Value) bool {
	return def != nil && def.Exists("values")
}

func ParseRecord(def *fastjson.Value) (RecordNode, error) {
	value, err := Parse(def.Get("values"))
	metadata := def.GetObject("metadata")
	return RecordNode{value, metadata}, err
}

func (node RecordNode) Metadata(key string) *fastjson.Value {
	if node.metadata == nil {
		return nil
	}

	return node.metadata.Get(key)
}
