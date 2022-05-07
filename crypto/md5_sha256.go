package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	data := []byte("richard")
	md5Count := int64(0)
	sha256Count := int64(0)

	go func() {
		for {
			//md5哈希计算
			md5.Sum(data)
			//原子计数+1
			atomic.AddInt64(&md5Count, 1)
		}
	}()

	go func() {
		for {
			sha256.Sum256(data)
			atomic.AddInt64(&sha256Count, 1)
		}
	}()

	time.Sleep(time.Second)
	fmt.Printf("md5Count:%d\n", md5Count)
	fmt.Printf("sha256Count:%d\n", sha256Count)
}