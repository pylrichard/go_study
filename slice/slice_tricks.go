package main

import "fmt"

//见https://github.com/golang/go/wiki/SliceTricks
var sa = make([]int, 6, 12)
//var sa = make([]int, 6)
var i, j = 2, 5
type callBack func(e int) bool

func main() {
	for i := 0; i < len(sa); i++ {
		sa[i] = i
	}
	fmt.Printf("%p, %v\n", sa, sa)

	//copy1()
	//copy2()

	//cut1()
	//cut2()

	//del1()
	//del2()
	//del3()
	//del4()
	//del5()

	//expandInternal()
	//expandTail()

	//n := filter1(sa, isEven)
	//sa = sa[:n]

	//even := filter2(sa, isEven)
	//fmt.Println(even)

	filter3(sa, isEven)

	//insert1()
	//insert2()

	//fmt.Println(popTail())
	//fmt.Println(popHead())

	//appendHead()

	/*
		sa在切片长度等于容量时，添加元素会生成新切片，见
		var sa = make([]int, 6)和insert1()

		sa在切片容量大于长度时，添加元素不会生成新切片，见
		var sa = make([]int, 6, 12)和expandTail()
	 */
	fmt.Printf("%p, %v\n", sa, sa)
}

/*
	复制，复制切片sa到切片sb
 */
func copy1() {
	//一次申请内存到位
	sb := make([]int, len(sa))
	copy(sb, sa)

	fmt.Println(sb)
}

func copy2() {
	//比copy()慢，但在复制之后有更多元素要添加，append()效率更高些
	sb := append([]int(nil), sa...)
	sc := append(sa[:0:0], sa...)

	fmt.Println(sb)
	fmt.Println(sc)
}

/*
	剪切，把切片sa中索引[i, j)的元素剪切
 */
func cut1() {
	sa = append(sa[:i], sa[j:]...)
}

/*
	删除，把切片sa中索引i的元素删除
 */
func del1() {
	sa = append(sa[:i], sa[i+1:]...)
}

func del2() {
	num := copy(sa[i:], sa[i+1:])
	//切片表达式截掉最后一个元素
	sa = sa[:i+num]
}

//只需要删除索引i的元素，无需保留切片元素原有的顺序
func del3() {
	//复制最后一个元素到索引i
	sa[i] = sa[len(sa)-1]
	//切片表达式截掉最后一个元素
	sa = sa[:len(sa)-1]
}

/*
	剪切或删除操作可能引起的内存泄露
	需要特别注意的是如果切片sa中元素是一个指针或包含指针字段的结构体(需要被垃圾回收)，剪切和删除会存在一个潜在的内存泄漏问题：
	一些具有值的元素仍被切片sa引用，因此无法被垃圾回收机制回收
 */
func cut2() {
	copy(sa[i:], sa[j:])
	for k, n := len(sa)-(j-i), len(sa); k < n; k++ {
		//赋值为nil或者类型零值
		sa[k] = 0
	}
	sa = sa[:len(sa)-(j-i)]
}

func del4() {
	copy(sa[i:], sa[i+1:])
	sa[len(sa)-1] = 0
	sa = sa[:len(sa)-1]
}

func del5() {
	//删除但不保留元素原有顺序
	sa[i] = sa[len(sa)-1]
	sa[len(sa)-1] = 0
	sa = sa[:len(sa)-1]
}

/*
	内部扩张
	在切片sa的索引i之后扩张j个元素
	使用两个append()完成，即先将索引i之后的元素追加到一个长度为j的切片sb后
	再将切片sc中的所有元素追加到切片a的索引i之后
	扩张这部分元素为类型零值
 */
func expandInternal() {
	sb := make([]int, j)
	sc := append(sb, sa[i:]...)
	sa = append(sa[:i], sc...)
}

/*
	尾部扩张
	将切片sa的尾部扩张j个元素的空间
 */
func expandTail() {
	sb := make([]int, j)
	sa = append(sa, sb...)
}

/*
	按照一定的规则将切片sa中的元素进行就地过滤
	假设过滤的条件已封装为keep()，使用for range遍历切片sa的所有元素，逐一调用keep()进行过滤
 */
func isEven(num int) bool {
	if num%2 == 0 {
		return true
	}

	return false
}

func filter1(s []int, cb callBack) int {
	n := 0
	for _, e := range s {
		if cb(e) {
			//保留元素
			s[n] = e
			n++
		}
	}
	//截取需要保留的元素
	s = s[:n]
	fmt.Printf("%p, %v\n", s, s)

	return n
}

func filter2(s []int, cb callBack) []int {
	var result []int
	for _, e := range s {
		if cb(e) {
			result = append(result, e)
		}
	}

	return result
}

func filter3(s []int, cb callBack) {
	//sb和切片sa共享相同的底层数组和容量
	sb := s[:0]
	fmt.Printf("%p, %v\n", sb, sb)
	for _, e := range s {
		if cb(e) {
			sb = append(sb, e)
		}
	}
	fmt.Printf("%p, %v\n", sb, sb)
	//处理必须被垃圾回收的元素
	for i := len(sb); i < len(s); i++ {
		s[i] = 0
	}
}

/*
	将新元素插入切片sa的索引i处
 */
func insert1() {
	//创建一个新切片，并将sa[i:]中的元素复制到该切片
	sb := append([]int{6}, sa[i:]...)
	sa = append(sa[:i], sb...)
}

//避免新切片的创建，以及由此产生的内存垃圾
func insert2() {
	//插入类型零值
	sa = append(sa, 0)
	fmt.Printf("%p, %v\n", sa, sa)
	copy(sa[i+1:], sa[i:])
	fmt.Printf("%p, %v\n", sa, sa)
	sa[i] = 6
}

/*
	弹出
 */

//将切片sa的最后一个元素弹出
func popTail() int {
	e := sa[len(sa)-1]
	sa = sa[:len(sa)-1]

	return e
}

//将切片sa的第一个元素弹出
func popHead() int {
	e := sa[0]
	sa = sa[1:]

	return e
}

//前插元素到切片sa头部
func appendHead() {
	sa = append([]int{-1}, sa...)
}

