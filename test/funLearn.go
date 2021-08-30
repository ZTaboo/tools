package test

import "fmt"

type Zero struct {
	Origin string
}

func (data *Zero) LearnFun1() {
	fmt.Println(data.Origin)
}

func LearnFun() {
	var zero = Zero{Origin: "hello"}
	zero.LearnFun1()
}
