package test

import (
	"fmt"
	"github.com/ZTaboo/tools/tool"
	"testing"
)

func TestRandomNum(t *testing.T) {
	res := tool.RandomNum()
	fmt.Println(res)
}
