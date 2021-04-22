package main

//见Go切片传递的隐藏危机.md
import "fmt"

func modifySlice1() {
	outerSlice := []string{"a", "a"}
	modifySlice1Impl(outerSlice)
	fmt.Print(outerSlice)
}

func modifySlice1Impl(innerSlice []string) {
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Println(innerSlice)
}

func modifySlice2() {
	outerSlice := []string{"a", "a"}
	modifySlice2Impl(&outerSlice)
	fmt.Print(outerSlice)
}

func modifySlice2Impl(innerSlice *[]string) {
	(*innerSlice)[0] = "b"
	(*innerSlice)[1] = "b"
	fmt.Println(*innerSlice)
}

func modifySlice3() {
	outerSlice := []string{"a", "a"}
	fmt.Printf("%p %v %p\n", &outerSlice, outerSlice, &outerSlice[0])
	modifySlice3Impl(outerSlice)
	fmt.Printf("%p %v %p\n", &outerSlice, outerSlice, &outerSlice[0])
}

func modifySlice3Impl(innerSlice []string) {
	fmt.Printf("%p %v %p\n", &innerSlice, innerSlice, &innerSlice[0])
	innerSlice = append(innerSlice, "a")
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	fmt.Printf("%p %v %p\n", &innerSlice, innerSlice, &innerSlice[0])
}

func modifySlice4() {
	outerSlice := []string{"a", "a"}
	modifySlice4Impl(&outerSlice)
	fmt.Print(outerSlice)
}

func modifySlice4Impl(innerSlice *[]string) {
	*innerSlice = append(*innerSlice, "a")
	(*innerSlice)[0] = "b"
	(*innerSlice)[1] = "b"
	fmt.Println(*innerSlice)
}

func modifySlice5() {
	outerSlice := []string{"a", "a"}
	modifySlice5Impl(outerSlice)
	fmt.Println(outerSlice)
}

func modifySlice5Impl(innerSlice []string) {
	innerSlice[0] = "b"
	innerSlice = append(innerSlice, "a")
	innerSlice[1] = "b"
	fmt.Println(innerSlice)
}

/*
	参考底层数组扩容1-4.png
 */
func modifySlice6() {
	outerSlice := make([]string, 0, 3)
	outerSlice = append(outerSlice, "a", "a")
	fmt.Printf("%p %v %p\n", &outerSlice, outerSlice, &outerSlice[0])
	modifySlice6Impl(outerSlice)
	//此时outerSlice的len为2
	fmt.Printf("len:%d cap:%d\n", len(outerSlice), cap(outerSlice))
	fmt.Printf("%p %v %p\n", &outerSlice, outerSlice, &outerSlice[0])
}

func modifySlice6Impl(innerSlice []string) {
	//此处切片没有扩容，指向同一个底层数组
	innerSlice = append(innerSlice, "a")
	fmt.Printf("%p %v %p\n", &innerSlice, innerSlice, &innerSlice[0])
	innerSlice[0] = "b"
	innerSlice[1] = "b"
	//此时innerSlice的len为3
	fmt.Printf("len:%d cap:%d\n", len(innerSlice), cap(innerSlice))
	fmt.Printf("%p %v %p\n", &innerSlice, innerSlice, &innerSlice[0])
}

func main() {
	//modifySlice1()
	//modifySlice2()
	//modifySlice3()
	//modifySlice4()
	//modifySlice5()
	modifySlice6()
}