package zero

import "fmt"

type Man struct {
	Name string
	Age  int
}

func (data *Man) ZeroFun() (string, int) {
	fmt.Println(data.Name, data.Age)
	return data.Name, data.Age
}
