package utils

import (
	"github.com/bwmarrin/snowflake"
	"reflect"
)

type IdGeneratorMap map[string]interface{}

const (
	Snow = "snow"
)

var idGeneratorMap IdGeneratorMap

func init() {
	idGeneratorMap = make(map[string]interface{})
	idGeneratorMap[Snow] = nextIdBySnowFlake
}

func NextId(keys ...string) interface{} {
	var key = Snow
	if keys != nil && len(keys) > 0 {
		key = keys[0]
	}
	return idGeneratorMap.Call(key)
}

func (m IdGeneratorMap) Call(key string, args ...interface{}) interface{} {
	f, ok := m[key]
	if !ok {
		return nil
	}

	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func {
		return nil
	}

	in := make([]reflect.Value, len(args))
	for i, arg := range args {
		in[i] = reflect.ValueOf(arg)
	}

	out := fv.Call(in)
	if len(out) == 0 {
		return nil
	}

	return out[0].Interface()
}

// NextIdBySnowFlake 雪花算法 https://github.com/bwmarrin/snowflake
func nextIdBySnowFlake() int64 {
	node, _ := snowflake.NewNode(1)
	id := node.Generate()

	return int64(id)
}
