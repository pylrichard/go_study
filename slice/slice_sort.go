package main

import (
	"fmt"
	"sort"
)

/*
	见https://tonybai.com/2020/11/26/slice-sort-in-go/
	见Go切片排序.md
 */
type IntSlice []int

func (s IntSlice) Len() int { return len(s) }

func (s IntSlice) Less(i, j int) bool { return s[i] < s[j] }

func (s IntSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type Lang struct {
	Name string
	Rank int
}

func main() {
	//sortSlice()
	//sortCustomTypeSlice()
	sortSliceWithSugar()
}

func sortSlice() {
	s := IntSlice([]int{89, 14, 8, 9, 17, 56, 95, 3})
	fmt.Println(s)
	sort.Sort(s)
	fmt.Println(s)
}

func sortCustomTypeSlice() {
	langs := []Lang {
		{"c", 2},
		{"go", 1},
		{"python", 3},
		{"lua", 4},
	}
	sort.Slice(langs, func(i, j int) bool {
		return langs[i].Rank < langs[j].Rank
	})
	fmt.Printf("%v\n", langs)
}

func sortSliceWithSugar() {
	s := []int{89, 14, 8, 9, 17, 56, 95, 3}
	fmt.Println(s)
	sort.Ints(s)
	fmt.Println(s)
}