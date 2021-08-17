package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, c)
	}
	for range os.Args[1:] {
		// output channel content
		fmt.Printf("%s \n", <-c)
	}
	fmt.Printf("total time: %.2fs\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	t := time.Now()
	r, err := http.Get(url) // 抛出网络错误
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	// size of r.Body
	written, err2 := io.Copy(ioutil.Discard, r.Body)
	r.Body.Close()
	if err2 != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // 抛出字符解析错误
	}
	f := time.Since(t).Seconds()
	// written -> ch
	ch <- fmt.Sprintf("%.2fs %7d %s", f, written, url)
}
