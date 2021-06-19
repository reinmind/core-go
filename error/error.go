package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	value, err := sqrt(-1)
	b := err != nil
	if b{
		fmt.Println(err)
	}
	fmt.Println(value)
}
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("math: negtive number")
	}
	return math.Sqrt(x), nil

}
