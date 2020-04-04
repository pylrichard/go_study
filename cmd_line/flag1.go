package main

import (
	"flag"
	"fmt"
)

// go run flag1.go -isOk -name test others
// go run flag1.go -help
// 定义待解析命令行参数，即以-开头的参数，-help不需要特别指定，可以自动处理
// 3个参数对应命令行参数名称，默认值，提示字符串
var isOk = flag.Bool("isOk", false, "bool param")
name := flag.String("name", "", "string param")

func main() {
	// flag定义完成后使用前调用
	flag.Parse()
	// 注意是参数变量指针
	fmt.Println("-isOk", *isOk)
	fmt.Println("-name", *name)
	// 命令行参数 指运行程序提供的参数
	// 已定义命令行参数 指通过flag.Var()形式预定义的参数
	// 未按照预定义解析的参数(non-flag)，通过flag.Args()获取，是一个字符串切片
	// 从第一个不能解析的参数开始，后面所有参数都无法解析。即使后面参数中含有预定义参数
	// go run flag1.go -isOk stop -name test others
	fmt.Println("other param:", flag.Args())
}