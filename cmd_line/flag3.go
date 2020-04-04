package main

import (
	"flag"
	"fmt"
	"strings"
)

type sliceValue []string

// 创建一个存储命令行参数的slice
func newSliceValue(values []string, ptr *[]string) *sliceValue {
	*ptr = values

	return (*sliceValue)(ptr)
}

/*
	type Value interface {
	    String() string
	    Set(string) error
	}
	实现flag包的Value接口，将命令行参数用逗号分隔存储到slice
*/
func (slice *sliceValue) Set(value string) error {
	*slice = sliceValue(strings.Split(value, ","))

	return nil
}

func (slice *sliceValue) Get() interface{} {
	return []string(*slice)
}

// 默认值default和返回值没关系
func (slice *sliceValue) String() string {
	return strings.Join([]string(*slice), ":")
	// *slice = sliceValue(strings.Split("default", ","))

	// return "It's none of my business"
}

// go run flag3.go -slice "go,java"
// flag对Duration非基本类型的支持使用类似方式
func main() {
	var lang []string
	flag.Var(newSliceValue([]string{}, &lang), "slice", "I like this lang")

	flag.Parse()
	fmt.Println(lang)
}