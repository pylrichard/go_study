package main

type X2 struct {
    a int
}
type Y2 struct {
    X2
    a int
}
type Z2 struct {
    Y2
    a int
}

func composite2() {
    x2 := X2{ a: 1 }
    y2 := Y2{
        X2: x2,
        a: 2,
    }
    z2 := Z2{
        Y2: y2,
        a: 3,
    }
    //z2.a, z2.Y2.a, z2.Y2.X2.a代表不同的字段
    println(z2.a, z2.Y2.a, z2.Y2.X2.a)
    z2 = Z2{}
    z2.a = 4
    z2.Y2.a = 5
    z2.Y2.X2.a = 6
    println(z2.a, z2.Y2.a, z2.Y2.X2.a)
}