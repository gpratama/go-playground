package main

import (
	"fmt"
	"math"
)

// Vertex of struct X, Y float64
type Vertex struct {
	X, Y float64
}

// Abs calculate the vertex length by using dot product against itself
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale the vertex by certain factor
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
