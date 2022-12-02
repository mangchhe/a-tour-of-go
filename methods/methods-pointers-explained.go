package main

import (
	"fmt"
)

type Vertex3 struct {
	X, Y float64
}

func Scale(v *Vertex3, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex3{3, 4}
	Scale(&v, 10)
	fmt.Println(v)
}
