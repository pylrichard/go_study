package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

/*
	见https://colobu.com/2016/04/01/how-to-get-goroutine-id/
	避免开发者滥用goroutine id实现goroutine local storage
	因为goroutine local storage很难进行垃圾回收
	在debug时log里的goroutine id是很好的一个监控信息
 */
func main() {
	fmt.Println("main", getID())
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		//没有这行代码，输出i都是10
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(i, getID())
		}()
	}
	wg.Wait()
}

func getID() int {
	var buf [64]byte
	//将当前的堆栈信息写入到一个slice，堆栈的第一行为goroutine ####
	//其中####就是当前的goroutine id
	//获取堆栈信息会影响性能，所以建议debug时才用它
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine"))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}

	return id
}
