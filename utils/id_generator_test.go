package utils

import (
	"fmt"
	"testing"
)

func TestNextIdBySnowFlake(t *testing.T) {
	id := NextIdBySnowFlake()
	fmt.Println(id)
}
