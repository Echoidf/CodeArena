package utils

import (
	"fmt"
	"testing"
)

func TestNextIdBySnowFlake(t *testing.T) {
	id := nextIdBySnowFlake()
	fmt.Println(id)
}

func TestNextId(t *testing.T) {
	id := NextId().(int64)
	fmt.Println(id)
}

func TestMd5Encode(t *testing.T) {
	encode := Md5Encode("absfdaf")
	fmt.Println(encode)
}
