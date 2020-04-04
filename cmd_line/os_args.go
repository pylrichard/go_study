package main

import (
	"strings"
	"fmt"
	"os"
	"strconv"
)

// go run os_args.go 1 3 -x
func main() {
	// os.Args类型是[]string字符串切片
	fmt.Println("len:", len(os.Args))
	fmt.Println(strings.Join(os.Args[1:], "\n"))
	fmt.Println(os.Args[1:])
	// os.Args第一个值是可执行文件的信息
	for i, arg := range os.Args[1:] {
		fmt.Println("arg " + strconv.Itoa(i) + ":", arg)
	}
}