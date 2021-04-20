package main

import (
	"fmt"
	"reflect"
)

func main() {
	slice()
	//array()
}


func slice() {
	s := []int{1, 2}
	fmt.Printf("%v, %T\n", s, s)
	fmt.Println(reflect.TypeOf(s), reflect.TypeOf(s).Kind())
	fmt.Printf("%p, %p, %p\n", s, s[0], s[1])
	//slice本身是一个指针，其地址与&s[0]地址相同
	fmt.Printf("%p, ,%p, %p, %p\n", &s, *(&s), &s[0], &s[1])
}

func array() {
	a := [2]int{1, 2}
	fmt.Printf("%v, %T\n", a, a)
	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(a).Kind())
	fmt.Printf("%p, %p, %p\n", a, a[0], a[1])
	//array本身不是一个指针，获取其地址要通过&a，&a地址与&a[0]地址相同
	fmt.Printf("%p, %p, %p\n", &a, &a[0], &a[1])
}