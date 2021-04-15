package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// 见http://junes.tech/2021/03/07/go-tip/go-tip-1/
func main() {
	simple()
	//useSignal()
	//one2One()
	//one2OneTwoChannel()
	//one2OneChanInChan()
	//one2OneUseContext()
	//one2Many()
	//one2ManyUseContext()
}

// curl http://localhost:8080/hello
// kill -9 {pid}
// 如果程序意外退出(例如panic)，无法和kill情况进行区分
func simple() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", mux)
}

func useSignal() {
	// 创建一个sig channel，捕获系统信号传递到sig中
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	// http服务改造成异步
	go http.ListenAndServe(":8080", mux)

	// 程序阻塞在这里，除非收到interrupt或者kill信号
	fmt.Println(<-sig)
}

// 既然子goroutine逻辑是有价值的，不想轻易丢失数据，那么为什么不把逻辑放到父goroutine中？其实很多时候都在滥用goroutine
// 抛开语言特性，从整体思考以下三个问题：
// 明确调用链路 - 梳理整个调用流程，区分关键和非关键的步骤，以及在对应步骤上发生错误时的处理方法
// 用MQ解耦服务 - 跨服务调用如果比较耗时，大部分时候更建议采用消息队列解耦
// 面向错误编程 - 关键业务的goroutine代码要考虑所有可能发生错误的点，保证程序退出或panic/recover不出现脏数据
func one2One() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	// 需要优雅退出
	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()

	fmt.Println(<-sig)
}

// 父goroutine通知子goroutine准备优雅退出
// 子goroutine通知父goroutine已经关闭完成
// 启用2个channel用来通信
func one2OneTwoChannel() {
	sig := make(chan os.Signal)
	stopCh := make(chan struct{})
	finishedCh := make(chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	go func(stopCh, finishedCh chan struct{}) {
		for {
			select {
			case <-stopCh:
				fmt.Println("stopped")
				finishedCh <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(stopCh, finishedCh)

	<-sig
	stopCh <- struct{}{}
	<-finishedCh
	fmt.Println("finished")
}

func one2OneChanInChan() {
	sig := make(chan os.Signal)
	stopCh := make(chan chan struct{})
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	go func(stopCh chan chan struct{}) {
		for {
			select {
			case ch := <-stopCh:
				fmt.Println("stopped")
				// 结束后通过ch通知父goroutine
				ch <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(stopCh)

	<-sig
	// ch作为一个channel，传递给子goroutine，待其结束后从中返回
	ch := make(chan struct{})
	stopCh <- ch
	// channel中传递的数据为error时，就能有更多信息
	<-ch
	fmt.Println("finished")
}

// context不仅可以传递数值，也可以控制子goroutine的生命周期
// context意思为上下文，最初的设计意为传递数值，就是一个数据流
// 而context又延伸出控制goroutine生命周期的功能，就成了控制流
// 这么看其实context角色不清晰
func one2OneUseContext() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
	ctx, cancel := context.WithCancel(context.Background())
	finishedCh := make(chan struct{})

	go func(ctx context.Context, finishedCh chan struct{}) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("stopped")
				finishedCh <- struct{}{}
				return
			default:
				time.Sleep(time.Second)
			}
		}
	}(ctx, finishedCh)

	<-sig
	cancel()
	<-finishedCh
	fmt.Println("finished")
}

// 服务端程序经常会处理各种逻辑，如操作数据库、读写文件、RPC调用等。根据其对原子性的要求，将处理逻辑区分为两种：
// 一种是无严格数据质量要求的，即程序直接崩溃也没有问题，比如普通查询
// 一种是有原子性要求，即不希望运行到一半就退出，比如写文件、修改数据等，最好是程序提供一定的缓冲时间，等待这部分逻辑处理完优雅退出
// 在复杂系统中为了保证数据质量，优雅退出是一个必要特性
func one2Many() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	// 10个goroutine模拟并发处理业务逻辑
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				// 这时捕获信号后10个goroutine就立刻停止
				// 希望程序能等当前周期休眠完，再优雅退出
				time.Sleep(time.Duration(i) * time.Second)
			}
		}(i)
	}

	fmt.Println(<-sig)
}

// 在追求优雅退出时要注意控制细粒度
// 比如一个http服务要控制整个http server的优雅退出
// 千万不要想着在main()控制每个http handler，即每个http请求的优雅退出，这样很难控制代码的复杂度
// 对于每个http请求的控制，应该交给http server框架去实现
// 所以main()中需要优雅退出的选项其实很有限
func one2ManyUseContext() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)
	ctx, cancel := context.WithCancel(context.Background())
	num := 10

	// 用wg控制多个子goroutine的生命周期
	wg := sync.WaitGroup{}
	wg.Add(num)

	for i := 0; i < num; i++ {
		go func(ctx context.Context) {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("stopped")
					return
				default:
					time.Sleep(time.Duration(i) * time.Second)
				}
			}
		}(ctx)
	}

	<-sig
	cancel()
	// 等待所有的子goroutine都优雅退出
	wg.Wait()
	fmt.Println("finished")
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}
