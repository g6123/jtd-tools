package parser

import (
	"github.com/g6123/jtd-tools/internal/util"
	"github.com/valyala/fastjson"
)

type DocumentNode struct {
	Root        Node
	Definitions map[string]DefinitionNode

	metadata *fastjson.Object
}

type DefinitionNode struct {
	Name  string
	Value Node
}

func ParseDocument(def *fastjson.Value) (DocumentNode, error) {
	node := DocumentNode{metadata: def.GetObject("metadata")}

	root, err := Parse(def)
	node.Root = root
	if err != nil {
		return node, err
	}

	child_defs := def.GetObject("definitions")
	if child_defs != nil {
		defs, err := ParseDocumentDefinitions(child_defs)
		node.Definitions = defs
		if err != nil {
			return node, err
		}
	}

	return node, nil
}

func ParseDocumentDefinitions(defs *fastjson.Object) (map[string]DefinitionNode, error) {
	results := make(chan util.Result[DefinitionNode])

	go func() {
		defs.Visit(func(key []byte, value_def *fastjson.Value) {
			value, err := Parse(value_def)

			results <- util.Result[DefinitionNode]{
				Value: DefinitionNode{
					Name:  string(key),
					Value: value,
				},
				Err: err,
			}
		})

		close(results)
	}()

	nodes := make(map[string]DefinitionNode, defs.Len())

	for result := range results {
		if result.Err != nil {
			return nodes, result.Err
		}

		nodes[result.Value.Name] = result.Value
	}

	return nodes, nil
}

func (node DocumentNode) Metadata(key string) *fastjson.Value {
	if node.metadata == nil {
		return nil
	}

	return node.metadata.Get(key)
}
