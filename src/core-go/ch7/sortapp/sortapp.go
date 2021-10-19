package main

import (
	"fmt"
	"sort"
)

type Point struct {
	x uint32
	y uint32
}
type PointArray []Point

func (pa PointArray) Len() int {
	return len(pa)
}
func (pa PointArray) Less(i, j int) bool {
	if pa[i].x < pa[j].x {
		return true
	} else if pa[i].x == pa[j].x && pa[i].y <= pa[j].y {
		return true
	}
	return false
}

func (pa PointArray) Swap(i, j int) {
	t := pa[i]
	pa[i] = pa[j]
	pa[j] = t
}

func main() {
	pa := PointArray{Point{1, 2}, Point{1, 3}, Point{2, 1}, {1, 1}}
	sort.Sort(pa)
	fmt.Println(pa)
}
