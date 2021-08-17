package main

import "fmt"

type tree struct {
	value          int
	lchild, rchild *tree
}

func AppendValues(values []int, t *tree) []int {
	if t != nil {
		values = AppendValues(values, t.lchild)
		values = append(values, t.value)
		values = AppendValues(values, t.rchild)
	}
	return values
}

func main() {
	t := tree{value: 10, lchild: &tree{value: 1}, rchild: &tree{value: 2}}
	fmt.Printf("node value: %d \t lchild: %v,rchild: %v \n", t.value, t.lchild, t.rchild)
	fmt.Printf("type of t: %T\t type of %T\n", t, &t)
}
