package main

import (
	"container/ring"
	"fmt"
)

func createRing(n int) *ring.Ring {
	r := ring.New(n)
	len := r.Len()
	for i := 0; i < len; i++ {
		r.Value = i
		r = r.Next()
	}

	return r
}

func printRing(r *ring.Ring) {
	r.Do(func (p interface{}) {
		fmt.Println(p.(int))
	})
}

func prev() {
	fmt.Println("Prev:")
	r := createRing(6)
	for i := 0; i < r.Len(); i++ {
		r = r.Prev()
		fmt.Println(r.Value)
	}
}

func move() {
	fmt.Println("Move:")
	r := createRing(6)
	// Move(n) moves n % r.Len() elements backward (n < 0) or forward (n >= 0)
	r = r.Move(-2)
	printRing(r)
}

func link() {
	fmt.Println("Link:")
	r1 := createRing(8)
	r2 := createRing(6)
	// Link connects ring r with ring s such that r.Next()
	// becomes s and returns the original value for r.Next()
	fmt.Println("diff ring")
	r3 := r1.Link(r2)
	fmt.Println("r1")
	printRing(r1)
	fmt.Println("r2")
	printRing(r2)
	fmt.Println("r3")
	printRing(r3)
	fmt.Println("same ring")
	r4 := r1.Move(3)
	r5 := r1.Link(r4)
	fmt.Println("r1")
	printRing(r1)
	fmt.Println("r4")
	printRing(r4)
	fmt.Println("r5")
	printRing(r5)
}

func unlink() {
	fmt.Println("Unlink:")
	r := createRing(6)
	// Unlink(n) removes n % r.Len() elements from the ring r
	// starting at r.Next(). If n % r.Len() == 0, r remains unchanged
	// The result is the removed subring
	r.Unlink(4)
	printRing(r)
}

func main() {
	link()
}