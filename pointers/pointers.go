package main

import (
	"fmt"
)

const SIZE = 12

func main() {
	arr := [SIZE]string{"Zeus", "Hera", "Poseidon", "Demeter", "Athena", "Apollo",
		"Ares", "Hephaestus", "Aphrodite", "Hermes", "Hestia", "Dionysus"}
	arr_ptr := &arr
	for i := 0; i < 12; i++ {
		fmt.Printf("%d: %s\n", (i + 1), arr_ptr[i])
	}
	if arr_ptr == &arr {
		fmt.Println("Dawn of Gods.")
	}

	var banner string = "hello pointers"
	fmt.Printf("address %x\n", &banner)
	success := make(chan bool)
	parseAddress(&banner)
	go parseValue(banner, success)
	<-success
}

/**
 * 值传递
 */
func parseValue(a string, success chan bool) {
	fmt.Printf("address %x", &a)
	fmt.Printf("value: %s\n", a)
	success <- true
}

/**
 * 引用传递
 */
func parseAddress(b *string) {
	*b = "hello value"
	fmt.Printf("address %x ", b)
	fmt.Printf("value: %s\n", *b)
}
