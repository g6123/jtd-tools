package parser

import (
	"github.com/g6123/jtd-tools/internal/util"
	"github.com/valyala/fastjson"
)

type ObjectNode struct {
	Properties []ObjectPropertyNode
	Nullable   bool

	metadata *fastjson.Object
}

type ObjectPropertyNode struct {
	Key        string
	IsOptional bool
	Value      Node
}

func IsObject(def *fastjson.Value) bool {
	return def != nil && (def.Exists("properties") || def.Exists("optionalProperties"))
}

func ParseObject(def *fastjson.Value) (ObjectNode, error) {
	node := ObjectNode{
		Nullable: def.GetBool("nullable"),
		metadata: def.GetObject("metadata"),
	}

	props_def := def.GetObject("properties")
	if props_def != nil {
		props, err := ParseObjectProperties(false, props_def)
		if err != nil {
			return node, err
		}

		node.Properties = append(node.Properties, props...)
	}

	opt_props_def := def.GetObject("optionalProperties")
	if opt_props_def != nil {
		props, err := ParseObjectProperties(true, opt_props_def)
		if err != nil {
			return node, err
		}

		node.Properties = append(node.Properties, props...)
	}

	return node, nil
}

func ParseObjectProperties(is_optional bool, def *fastjson.Object) ([]ObjectPropertyNode, error) {
	results := make(chan util.Result[ObjectPropertyNode])

	go func() {
		def.Visit(func(key []byte, prop_def *fastjson.Value) {
			prop, err := ParseObjectProperty(string(key), is_optional, prop_def)

			results <- util.Result[ObjectPropertyNode]{
				Value: prop,
				Err:   err,
			}
		})

		close(results)
	}()

	props := make([]ObjectPropertyNode, 0, def.Len())

	for result := range results {
		if result.Err != nil {
			return props, result.Err
		}

		props = append(props, result.Value)
	}

	return props, nil
}

func ParseObjectProperty(key string, is_optional bool, def *fastjson.Value) (ObjectPropertyNode, error) {
	node := ObjectPropertyNode{Key: key, IsOptional: is_optional}
	node_value, err := Parse(def)
	node.Value = node_value
	return node, err
}

func (node ObjectNode) Metadata(key string) *fastjson.Value {
	if node.metadata == nil {
		return nil
	}

	return node.metadata.Get(key)
}
