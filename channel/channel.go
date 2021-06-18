package main

import "fmt"

func main() {
	c := make(chan int, 10)
	fib(10, c)
	for i := range c {
		fmt.Println(i)
	}
}

/**
 *无情的并发机器
 */
func fib(n int, c chan int) {
	x, y := 1, 1
	for i := 1; i <= n; i++ {
		c <- x
		z := y
		y, x = x+y, z
	}
	close(c)
}
