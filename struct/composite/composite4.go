package main

type X4 struct {
    a int
}

type Y4 struct {
    X4
}

type Z4 struct {
    *X4
}

func (x4 X4) Get() int {
    return x4.a
}

func (x4 *X4) Set(i int) {
    x4.a = i
}

func composite4() {
    x4 := X4{ a: 1 }
    y4 := Y4{
        X4: x4,
    }
    println(y4.Get())
    //此处编译器做了自动转换
    y4.Set(2)
    println(y4.Get())
    //为了不让编译器做自动转换，使用方法表达式调用方式
    //Y4内嵌字段X4，所以type y4的方法集是Get()，type *Y4的方法集是Set()/Get()
    (*Y4).Set(&y4, 3)
    //type y4的方法集并没有Set()，所以下一句编译不能通过
    //Y.Set(y, 3)
    println(y4.Get())
    z4 := Z4{
        X4: &x4,
    }
    //按照嵌套字段的方法集的规则
    //Z4内嵌字段＊X4 ，所以type Z4和type *Z4方法集都包含类型X4定义的方法Get()和Set()
    //为了不让编译器做自动转换，仍然使用方法表达式调用方式
    Z4.Set(z4, 4)
    println(z4.Get())
    (*Z4).Set(&z4, 5)
    println(z4.Get())
}