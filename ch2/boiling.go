package main

import "fmt"

const boilingF = 211.0

func Main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g *F or %g *C\n", f, c)
}