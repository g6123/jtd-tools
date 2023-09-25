package schema

import "github.com/g6123/jtd-tools/internal/parser"

type RecordSchema struct {
	Type                 string `json:"type"`
	AdditionalProperties Schema `json:"additionalProperties"`
}

func FromRecord(node parser.RecordNode) RecordSchema {
	return RecordSchema{
		Type:                 "object",
		AdditionalProperties: From(node.Value),
	}
}
