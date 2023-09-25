package parser

import (
	"github.com/g6123/jtd-tools/internal/util"
	"github.com/valyala/fastjson"
)

type DiscriminatedUnionNode struct {
	Disciriminator string
	Mapping        map[string]DiscriminatedVariantNode
	Nullable       bool

	metadata *fastjson.Object
}

type DiscriminatedVariantNode struct {
	Name  string
	Value ObjectNode
}

func IsDiscriminatedUnion(def *fastjson.Value) bool {
	return def != nil && def.Exists("discriminator") && def.Exists("mapping")
}

func ParseDiscriminatedUnion(def *fastjson.Value) (DiscriminatedUnionNode, error) {
	node := DiscriminatedUnionNode{
		Disciriminator: string(def.GetStringBytes("discriminator")),
		Nullable:       def.GetBool("nullable"),
		metadata:       def.GetObject("metadata"),
	}

	mapping, err := ParseDiscriminatedMapping(def.GetObject("mapping"))
	if err != nil {
		return node, err
	}

	node.Mapping = mapping
	return node, nil
}

func ParseDiscriminatedMapping(def *fastjson.Object) (map[string]DiscriminatedVariantNode, error) {
	results := make(chan util.Result[DiscriminatedVariantNode])

	go func() {
		def.Visit(func(key []byte, value_def *fastjson.Value) {
			value, err := ParseObject(value_def)

			results <- util.Result[DiscriminatedVariantNode]{
				Value: DiscriminatedVariantNode{
					Name:  string(key),
					Value: value,
				},
				Err: err,
			}
		})

		close(results)
	}()

	nodes := make(map[string]DiscriminatedVariantNode, def.Len())

	for result := range results {
		if result.Err != nil {
			return nil, result.Err
		}

		nodes[result.Value.Name] = result.Value
	}

	return nodes, nil
}

func (node DiscriminatedUnionNode) Metadata(key string) *fastjson.Value {
	if node.metadata == nil {
		return nil
	}

	return node.metadata.Get(key)
}
