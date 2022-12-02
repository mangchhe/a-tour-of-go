package main

import "fmt"

type I2 interface {
	M()
}

type T2 struct {
	S string
}

func (t *T2) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I2

	var t *T2
	i = t
	describe(i)
	i.M()

	i = &T2{"hello"}
	describe(i)
	i.M()
}

func describe(i I2) {
	fmt.Printf("(%v, %T)\n", i, i)
}
