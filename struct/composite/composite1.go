package main

import "fmt"

type X1 struct {
	a int
}

type Y1 struct {
	X1
	b int
}

type Z1 struct {
	Y1
	c int
}

func composite1() {
	x1 := X1{ a:1 }
	y1 := Y1{
		X1: x1,
		b: 2,
	}
	z1 := Z1{
		Y1: y1,
		c: 3,
	}
	//z1.a, z1.Y1.a, z1.Y1.X1.a三者是等价的，z1.a/z1.Y1.a是z1.Y1.X1.a的简写
	fmt.Println(z1.a, z1.Y1.a, z1.Y1.X1.a)
	z1 = Z1{}
	z1.a = 2
	fmt.Println(z1.a, z1.Y1.a, z1.Y1.X1.a)
}