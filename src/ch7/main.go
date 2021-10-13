package main

import "fmt"

type Point struct {
	x float64
	y float64
}

// 实现了String() string接口的结构体都能调用Printf,Println函数
func (p Point) String() string {
	return fmt.Sprintf("{x:%.2f,y:%.2f}", p.x, p.y)
}
func main() {
	point := Point{3, 4}
	fmt.Println(point)
}
