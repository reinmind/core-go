package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// 简单的获取http信息的程序
func main() {
	for _, url := range os.Args[1:] {
		response, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch %v \n", err)
			os.Exit(1)
		}
		// ReadAll接受一个ioReader参数，将reader中的内容变成字符串
		b, err2 := ioutil.ReadAll(response.Body)
		response.Body.Close()
		if err2 != nil {
			fmt.Fprintf(os.Stderr, "%v \n", err2)
			os.Exit(1)
		}
		// 通过ioutil获取response body
		fmt.Printf("%s", b)
	}
}
