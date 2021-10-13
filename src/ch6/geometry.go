package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Point type method
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}
func main() {
	p := Point{0, 0}
	q := Point{3, 4}
	q.ScaleBy(0.5)
	fmt.Println(p.Distance(q))
	fmt.Println(Distance(p, q))
}
