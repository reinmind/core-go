package main

// fmt 函数会自动排序
import (
	"fmt"
	"os"
)

func Main() {
	var s, sep string
	// os.Args接受参数
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "/"
	}
	fmt.Println(s)

}
