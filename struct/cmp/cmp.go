package cmp

import (
	"fmt"
	"reflect"
)

type Value1 struct {
	Name string
	Gender string
}

type Value2 struct {
	Name string
	Gender *string
}

type Value3 struct {
	Name string
	GoodAt []string
}

type Value4 struct {
	Name string
}

type Value5 struct {
	Name string
}

func Cmp1() {
	v1 := Value1{ Name: "pyl", Gender: "male" }
	v2 := Value1{ Name: "pyl", Gender: "male" }
	if v1 == v2 {
		fmt.Println("true")
		return
	}
	fmt.Println("false")
}

func Cmp2() {
	v1 := Value2{ Name: "pyl", Gender: new(string) }
    v2 := Value2{ Name: "pyl", Gender: new(string) }
    if v1 == v2 {
        fmt.Println("true")
        return
    }

    fmt.Println("false")
}

func Cmp3() {
	// v1 := Value3{ Name: "pyl", GoodAt: []string{"math", "chinese", "english"}}
    // v2 := Value3{ Name: "pyl", GoodAt: []string{"math", "chinese", "english"}}
    // if v1 == v2 {
    //     fmt.Println("true")
    //     return
    // }

    fmt.Println("false")
}

func Cmp4() {
	v1 := Value4{ Name: "pyl" }
	v2 := Value5{ Name: "pyl" }
	// if v1 == v2 {
    //     fmt.Println("true")
    //     return
    // }
	//强制转换可行
	//不可比较类型依然不行
	if v1 == Value4(v2) {
        fmt.Println("true")
        return
	}

    fmt.Println("false")
}

func Cmp5() {
	gender := new(string)
	v1 := Value2 {Name: "pyl", Gender: gender }
    v2 := Value2 {Name: "pyl", Gender: gender }
	if v1 == v2 {
        fmt.Println("true")
        return
    }

    fmt.Println("false")
}

func Cmp6() {
	v1 := Value3{ Name: "pyl", GoodAt: []string{"math", "chinese", "english"}}
    v2 := Value3{ Name: "pyl", GoodAt: []string{"math", "chinese", "english"}}
    if reflect.DeepEqual(v1, v2) {
        fmt.Println("true")
        return
    }

    fmt.Println("false")
}