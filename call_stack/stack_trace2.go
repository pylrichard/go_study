package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"runtime/debug"
	"runtime/pprof"
	"syscall"
	"time"
)

//见https://colobu.com/2016/12/21/how-to-dump-goroutine-stack-traces/
func main() {
	//对应fn2()调用fn4()
	setupSigUsr1Trap()

	go logic()
	fn1()
}

func fn1() {
	fn2()
}

func fn2() {
	//fn3()
	fn4()
	//fn5()
	//fn6()
}

//panic输出当前goroutine的stack trace
//当前goroutine的状态是running
//输出所有goroutine堆栈信息，设置GOTRACEBACK=1
//GOTRACEBACK=1 go run stack_trace2.go
//https://golang.org/pkg/runtime/有介绍此环境变量
func fn3() {
	panic("panic from fn3")
}

//程序没有发生panic，但是程序有问题，假死不工作
//想看哪儿出现了问题，可以给程序发送SIGQUIT信号，输出stack trace信息
//kill -SIGQUIT <pid>
func fn4() {
	time.Sleep(time.Hour)
}

//上面的情况是必须要求程序退出才能打印stack trace信息
//但是有时候只需要跟踪分析程序问题，而不希望程序中断运行
//可以通过一个API或者监听一个信号，然后调用相应的方法输出当前goroutine的stack trace
func fn5() {
	debug.PrintStack()
	time.Sleep(time.Hour)
}

//输出所有goroutine的stack trace
func fn6() {
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	time.Sleep(time.Hour)
}

//此goroutine的状态是sleep
func logic() {
	time.Sleep(time.Hour)
}

//使用runtime.Stack()得到所有goroutine的stack trace信息
//为了更方便随时得到所有goroutine的stack trace信息，可以监听SIGUSER1信号
//通过kill -SIGUSER1 <pid>发送信号，不必担心kill会将程序杀死
//配置http/pprof可以通过下面的地址访问所有的goroutine的stack trace
//http://localhost:8888/debug/pprof/goroutine?debug=2
func setupSigUsr1Trap() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGUSR1)
	go func() {
		for range c {
			dumpStacks()
		}
	}()
}

//见https://github.com/moby/moby/blob/95fcf76cc64a4acf95c168e8d8607e3acf405c13/pkg/signal/trap.go
//对比https://github.com/moby/moby/blob/master/pkg/signal/trap.go
func dumpStacks() {
	buf := make([]byte, 16384)
	buf = buf[:runtime.Stack(buf, true)]
	fmt.Println("begin goroutine stack dump...")
	fmt.Printf("%s\n", buf)
	fmt.Println("end goroutine stack dump...")
}