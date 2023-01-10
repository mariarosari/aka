package main

import (
	"fmt"
)

func main() {
	var TabInt [1000]int
	var N, x, i int
	N = 1000
	for i = 0; i < 1000; i++ {
		TabInt[i] = i + 1
	}
	fmt.Print("Data yang ingin dicari: ")
	fmt.Scanln(&x)
	fmt.Println(seqSearch(TabInt, N, x))

}

func seqSearch(T [1000]int, N int, x int) bool {
	var ketemu bool = false
	var k int = 0
	for !ketemu && k < N {
		ketemu = T[k] == x
		k++
	}
	return ketemu
}
