package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

type Point struct {
	X float32
	Y float32
}

func (p *Point) MoveXForward() {
	p.X += 1
}
func (p *Point) MoveYForward() {
	p.Y += 1
}
func (p *Point) MoveForward() {
	p.X += 1
	p.Y += 1
}
func (p Point) MoveBackward() {
	p.X -= 1
	p.Y -= 1
}
func (p Point) String() string {
	return fmt.Sprintf("X:%.2f , Y:%.2f", p.X, p.Y)
}

func main() {

	p := Point{X: 0, Y: 0}
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		// handle one connection concurrently
		go handleConn(conn, &p)
	}
}

// communicate by sharing memory causing concurrent problems
func handleConn(c net.Conn, p *Point) {

	defer c.Close()
	for {
		p.MoveXForward()
		time.Sleep(20 * time.Millisecond)
		p.MoveYForward()
		time.Sleep(20 * time.Millisecond)
		p.MoveXForward()
		time.Sleep(20 * time.Millisecond)
		p.MoveYForward()
		_, err := io.WriteString(c, time.Now().Format("15:04:05  "))
		io.WriteString(c, fmt.Sprintf("%v\n", *p))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
