package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	message, err := greetings.Hello("")

	log.SetPrefix("greetings: ")
	log.SetFlags(log.LstdFlags)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
