package main

import "fmt"

type X3 struct {
    a int
}

type Y3 struct {
    X3
    b int
}

type Z3 struct {
    Y3
    c int
}

func (x3 X3) Print() {
    fmt.Printf("In X3, a ＝ %d\n", x3.a)
}

func (x3 X3) XPrint() {
    fmt.Printf("In X3, a ＝ %d\n", x3.a)
}

func (y3 Y3) Print() {
    fmt.Printf("In Y3, b ＝ %d\n", y3.b)
}

func (z3 Z3) Print() {
    fmt.Printf("In Z3, c ＝ %d\n", z3.c)
    //显式使用完全路径调用内嵌字段的方法
    z3.Y3.Print()
    z3.Y3.X3.Print()
}

func composite3() {
    x3 := X3{ a: 1 }
    y3 := Y3{
        X3: x3,
        b: 2,
    }
    z3 := Z3{
        Y3: y3,
        c: 3,
    }
    //从外向内查找，首先找到的是z3的Print()
    z3.Print()
    //从外向内查找，最后找到的是x3的XPrint()
    z3.XPrint()
    z3.Y3.XPrint()
}