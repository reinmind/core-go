package main

import (
	"fmt"
	"reflect"
)

type X struct {
	str string
}

func main() {
	x := "hello world"

	t := reflect.TypeOf(x)
	z := X{"world hello"}
	fmt.Println(t)
	fmt.Println(reflect.TypeOf(z))
	fmt.Println(reflect.ValueOf(z))
}
