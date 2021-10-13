package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

var gv uint64

func main() {
	gv = 1
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
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	x := 1
	defer c.Close()
	for {
		x += 1
		gv += 1
		_, err := io.WriteString(c, time.Now().Format("15:04:05  "))
		io.WriteString(c, fmt.Sprintf("lv:%d gv:%d\n", x, gv))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
