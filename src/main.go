package main

import (
	"fmt"
	"prac/jwt"
)

func main() {
	simple := jwt.GenSimple()
	fmt.Printf("%s\n%s\n",simple,jwt.Decode(simple))
}
