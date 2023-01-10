package main

import "fmt"

func main() {
	var TabInt [9999]int
	var N, x, i int
	N = 1000
	for i = 0; i < 1000; i++ {
		TabInt[i] = i + 1
	}
	fmt.Print("Data yang ingin dicari: ")
	fmt.Scanln(&x)
	fmt.Println(binSearch(TabInt, N, x))
}

func binSearch(tab [9999]int, n int, x int) bool {
	var left, right, mid int
	left = 1
	right = n
	mid = (left + right) / 2
	for left <= right && tab[mid] != x {
		if x < tab[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	return mid > 0 && tab[mid] == x
}
