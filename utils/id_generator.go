package utils

import (
	"github.com/bwmarrin/snowflake"
)

var IdGeneratorMap map[string]func() string

func init() {
	IdGeneratorMap = make(map[string]func() string)
}

// NextIdBySnowFlake 雪花算法 https://github.com/bwmarrin/snowflake
func NextIdBySnowFlake() int64 {
	node, _ := snowflake.NewNode(1)
	id := node.Generate()

	return int64(id)
}
