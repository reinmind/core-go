package main

import "fmt"

type Book struct {
	title   string
	author  string
	subject string
	id      uint32
}

func main() {
	book1 := Book{"Harry Potter", "J.K Rowling", "novel", 1}
	book2 := Book{id: 2, title: "活着", author: "余华", subject: "现实主义"}
	fmt.Println(book1)
	fmt.Printf("address: %x\n", &book2)
	infoBook(&book2)
	copyInfo(book2)
}

func infoBook(book *Book) {
	fmt.Printf("address: %x", book)
	fmt.Printf("title:%s subject:%s author:%s\n", (*book).title, book.subject, (*book).author)
}

func copyInfo(book Book) {
	fmt.Printf("address: %x", &book)
	fmt.Printf("title:%s subject:%s author:%s\n", book.title, book.subject, book.author)
}
