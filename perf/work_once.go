package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

func counter() {
	s := make([]int, 0)
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 15
		s = append(s, c)
	}
}

/*
	counter()占用cpu时间过多，实际是append()中内存重新分配造成的，
	简单做法就是事先申请一个大内存，避免频繁进行内存分配
 */
func betterCounter() {
	s := [100000]int{0}
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 15
		s[i] = c
	}
}

func workOnce(wg *sync.WaitGroup) {
	betterCounter()
	//counter()
	wg.Done()
}

func WorkOnce() {
	var cpuProfile = flag.String("cpu_profile", "", "write cpu profile to file")
	var memProfile = flag.String("mem_profile", "", "write mem profile to file")
	flag.Parse()
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		err = pprof.StartCPUProfile(f)
		if err != nil {
			log.Fatal(err)
		}
		defer pprof.StopCPUProfile()
	}

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go workOnce(&wg)
	}
	wg.Wait()

	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal(err)
		}
		err = pprof.WriteHeapProfile(f)
		if err != nil {
			log.Fatal(err)
		}
		_ = f.Close()
	}
}