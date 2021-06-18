package main

import "fmt"

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
}
