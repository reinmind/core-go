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
func (p Point) ScaleBy2(factor float64) {
	p.X *= factor
	p.Y *= factor
}
func Move(p Point, x float64, y float64) {
	p.X = x
	p.Y = y
}
func Move2(p *Point, x float64, y float64) {
	p.X = x
	p.Y = y
}
func main() {
	p := Point{0, 0}
	q := Point{3, 4}
	// 改变原来的值
	q.ScaleBy(0.5)
	// 不改变原来的值
	q.ScaleBy2(2)
	// 没有改变q的值
	Move(q, 3, 4)
	// 改变了q的值
	Move2(&q, 3, 4)
	fmt.Println(p.Distance(q))
	fmt.Println(Distance(p, q))
}
