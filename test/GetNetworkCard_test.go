package test

import (
	"fmt"
	"github.com/ZTaboo/tools/tool"
	"testing"
)

func TestGetNetworkCard(t *testing.T) {
	res := tool.GetNetworkCard()
	fmt.Println(res)
}
