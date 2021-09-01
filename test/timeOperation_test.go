package test

import (
	"fmt"
	"github.com/ZTaboo/tools/tool"
	"testing"
)

func TestTimeOperation(t *testing.T) {
	var module tool.ParseTime
	module.DataFormat = tool.Time5
	module.StrData = "20210902"
	fmt.Println(module.ParseData())
}
