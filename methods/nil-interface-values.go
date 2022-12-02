package main

import "fmt"

type I3 interface {
	M()
}

func main() {
	var i I
	describe2(i)
	i.M()
}

func describe2(i I3) {
	fmt.Printf("(%v, %T)\n", i, i)
}
