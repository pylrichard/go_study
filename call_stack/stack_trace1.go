package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
)

//见https://colobu.com/2018/11/03/get-function-name-in-go/
func main() {
	biz()
}

func biz() {
	fmt.Println(printMyName(), printCallerName())
	bar()
}

func bar() {
	fmt.Println(printMyName(), printCallerName())
	//trace()
	//traceMore()

	//程序中遇到Error，但是不期望程序panic，只是想把堆栈信息打印出来以便跟踪调试
	debug.PrintStack()
}

func printMyName() string {
	return printFuncName(2)
}

func printCallerName() string {
	return printFuncName(3)
}

func printFuncName(skip int) string {
	//Caller()返回函数调用栈的某一层的程序计数器、文件信息、行号
	//0代表当前函数，也是调用runtime.Caller()的函数。1代表上一层调用者
	pc, _, _, _ := runtime.Caller(skip)

	return runtime.FuncForPC(pc).Name()
}

func trace() {
	pc := make([]uintptr, 10)
	//Callers返回调用栈的程序计数器
	//0代表Callers()本身，和Caller()的参数意义不一样，历史原因造成的。1才对应上面的0
	n := runtime.Callers(0, pc)
	for i := 0; i < n; i++ {
		f := runtime.FuncForPC(pc[i])
		file, line := f.FileLine(pc[i])
		fmt.Printf("%s:%d %s\n", file, line, f.Name())
	}
}

func traceMore() {
	pc := make([]uintptr, 10)
	n := runtime.Callers(0, pc)
	//获得整个栈的信息，可以使用CallersFrames()，省去遍历调用FuncForPC()
	frames := runtime.CallersFrames(pc[:n])
	for {
		frame, more := frames.Next()
		fmt.Printf("%s:%d %s\n", frame.File, frame.Line, frame.Function)
		if !more {
			break
		}
	}
}