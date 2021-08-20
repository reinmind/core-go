package itf

import (
	"io"
	"os"
)

func Main() {
	var w io.Writer = os.Stdout
	w.Write([]byte("hello world\n"))
	//n, err := w.Write([]byte("hello world"))
	// if err == nil {
	// 	fmt.Printf("n: %v\n", n)
	// }
}
