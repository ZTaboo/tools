package test

import (
	"fmt"
	"github.com/ZTaboo/tools/tool"
	"testing"
)

func TestTimeOperation(t *testing.T) {
	var module tool.ParseTime
	module.DataFormat = tool.Time5
	module.StrData = "2021-08-01"
	fmt.Println(module.ParseData())
}
