package ch4

import "fmt"

type Tree struct {
	value          int
	lchild, rchild *Tree
}

func AppendValues(values []int, t *Tree) []int {
	if t != nil {
		values = AppendValues(values, t.lchild)
		values = append(values, t.value)
		values = AppendValues(values, t.rchild)
	}
	return values
}

func TreeMain() {
	t := Tree{value: 10, lchild: &Tree{value: 1}, rchild: &Tree{value: 2}}
	fmt.Printf("node value: %d \t lchild: %v,rchild: %v \n", t.value, t.lchild, t.rchild)
	fmt.Printf("type of t: %T\t type of %T\n", t, &t)
}
