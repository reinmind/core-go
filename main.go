package main

import (
	"fmt"
	"prac/jwt"
)

func main() {
	s := jwt.GenSimple()
	fmt.Printf("%s\n", s)
	jwt.Decode(s)
}
