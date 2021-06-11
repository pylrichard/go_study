package main

import (
	"flag"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func anotherCounter1(wg *sync.WaitGroup) {
	wg.Done()
	s := []int{0}
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 15
		s = append(s, c)
	}
}

func anotherCounter2(wg *sync.WaitGroup) {
	wg.Done()
	s := [100000]int{0}
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 15
		s[i] = c
	}
}

func anotherCounter3(wg *sync.WaitGroup, m *sync.Mutex) {
	wg.Done()
	s := [100000]int{0}
	c := 1
	for i := 0; i < 100000; i++ {
		m.Lock()
		c = i + 15
		s[i] = c
		m.Unlock()
	}
}

func Trace() {
	//4核运行程序
	runtime.GOMAXPROCS(4)
	var traceProfile = flag.String("trace_profile", "", "write trace profile to file")
	flag.Parse()

	if *traceProfile != "" {
		f, err := os.Create(*traceProfile)
		if err != nil {
			log.Fatal(err)
		}
		_ = trace.Start(f)
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				log.Fatal(err)
			}
		}(f)
		defer trace.Stop()
	}

	var wg sync.WaitGroup
	var m sync.Mutex
	num := 3
	wg.Add(num)
	for i := 0; i < num; i++ {
		m.Lock()
		go anotherCounter2(&wg)
		time.Sleep(time.Millisecond)
		m.Unlock()
	}
	wg.Wait()
	time.Sleep(800 * time.Microsecond)
}