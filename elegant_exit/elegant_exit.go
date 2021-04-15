package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	simple()
	useSignal()
	one2One()
	one2Many()
	twoChannel()
}

func simple() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", mux)
}

func useSignal() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)
	go http.ListenAndServe(":8080", mux)

	fmt.Println(<-sig)
}

func one2One()  {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				time.Sleep(time.Duration(i) * time.Second)
			}
		}(i)
	}

	fmt.Println(<-sig)
}

func twoChannel() {
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

func one2Many() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL)

	go func() {
		for {
			time.Sleep(time.Second)
		}
	}()

	fmt.Println(<-sig)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello")
}