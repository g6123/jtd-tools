package parser

import "github.com/valyala/fastjson"

type EnumNode struct {
	Cases    []string
	Nullable bool

	metadata *fastjson.Object
}

func IsEnum(def *fastjson.Value) bool {
	return def != nil && def.Exists("enum")
}

func ParseEnum(def *fastjson.Value) (EnumNode, error) {
	node := EnumNode{
		Nullable: def.GetBool("nullable"),
		metadata: def.GetObject("metadata"),
	}

	cases_def := def.GetArray("enum")
	node.Cases = make([]string, len(cases_def))

	for i, case_def := range cases_def {
		case_value, err := case_def.StringBytes()
		if err != nil {
			return node, err
		}

		node.Cases[i] = string(case_value)
	}

	return node, nil
}

func (node EnumNode) Metadata(key string) *fastjson.Value {
	if node.metadata == nil {
		return nil
	}

	return node.metadata.Get(key)
}
