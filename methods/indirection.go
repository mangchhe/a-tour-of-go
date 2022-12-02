package main

import (
	"fmt"
	"math"
)

type Vertex4 struct {
	X, Y float64
}

func (v *Vertex4) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex4, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex4) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex4) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex4{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex4{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))

	fmt.Println(v, p)
}
