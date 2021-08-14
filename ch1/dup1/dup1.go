package main

// 第一个版本的dup输出标准输入流中的出现多次的行，
// 在行内容前是出现次数的计数。
// 这个程序将引入if 表达式，map内置数据结构和bufio的package。
import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map是Go语言内置的key/value型数据结构，这个数据结构能够提供常数时间的存取操作
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// 不用担心map在未初始化某个key时就去对其进行++操作，Go语言在碰到这种情况时，会自动将其初始化 为0，然后再进行操作。
		counts[input.Text()]++
	}
	// range会返回两个值， 一个key和在map对应这个key的value
	// 对map进行range循环时，其迭代顺序是不确定的，从实践来看， 很可能每次运行都会有不一样的结果
	// entry : entries
	for line, n := range counts {
		if n > 1 {
			/**
			verbs
			%q 带双引号的字符串 "abc" 或 带单引号的 rune 'c'
			%v 会将任意变量以易读的形式打印出来
			%T 打印变量的类型
			*/
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
