package main

import (
    "fmt"
    "flag"
)

// Hello flag.Value接口
type Hello string

// Set 实现flag.Value接口，要求receiver是指针
func (hello *Hello) Set(str string) error {
	value := fmt.Sprintf("Hello %v", str)
	*hello = Hello(value)

	return nil
}

func (hello *Hello) String() string {
	return fmt.Sprintf("%f", *hello)
}

// go run flag4.go -hello pylrichard go java python
func main() {
	var hello Hello
	flag.Var(&hello, "hello", "hello param")

	flag.Parse()
	others := flag.Args()
	fmt.Println("hello:", hello)
	fmt.Println("other:", others)
}