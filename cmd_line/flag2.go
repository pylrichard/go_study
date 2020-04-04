package main

import (
	"os"
	"fmt"
	"flag"
)

var (
	flagValue int
	help *bool
)

func init() {
	flag.IntVar(&flagValue, "flagValue", 1234, "help for flagValue")
	// 返回一个相应类型的指针
	help = flag.Bool("h", false, "this help")
	flag.Usage = usage
}

// go run flag2.go -flagValue
// go run flag2.go -h
// go build -o flag2 && ./flag2 -h
func main() {
	flag.Parse()
	// NArg()获得non-flag的个数
	for i := 0; i != flag.NArg(); i++ {
		fmt.Println("arg[%d] = %s\n", i, flag.Arg(i))
	}
	fmt.Println(flag.NFlag())
	fmt.Println(flagValue)
	if *help {
		flag.Usage()
	}
}

func usage() {
	// 注意此处的``和`)，否则显示格式不对
	fmt.Fprintf(os.Stderr, `
Usage: flag1.go -flagValue -h
Options:
`)
	flag.PrintDefaults()
}