package parser

import (
	"fmt"

	"github.com/valyala/fastjson"
)

type Node interface {
	Metadata(key string) *fastjson.Value
}

func Parse(def *fastjson.Value) (Node, error) {
	if IsObject(def) {
		return ParseObject(def)
	}

	if IsDiscriminatedUnion(def) {
		return ParseDiscriminatedUnion(def)
	}

	if IsRecord(def) {
		return ParseRecord(def)
	}

	if IsArray(def) {
		return ParseArray(def)
	}

	if IsEnum(def) {
		return ParseEnum(def)
	}

	if IsRef(def) {
		return ParseRef(def)
	}

	if IsPrimitive(def) {
		return ParsePrimitive(def)
	}

	if IsAny(def) {
		return ParseAny(def)
	}

	return AnyNode{}, fmt.Errorf("unkown node")
}
