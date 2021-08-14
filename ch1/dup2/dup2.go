package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args
	if len(files) < 1 {
		// 标准输入作为流
		// file /dev/stdin
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files[1:] {
			f, err := os.Open(arg)
			if err != nil {
				//
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n < 2 {
			continue
		}
		fmt.Printf("%d\t%s\n", n, line)
	}
}

// 这里的counts实际上是map的指针副本,虽然值不一样,但是地址指向是一样的
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
