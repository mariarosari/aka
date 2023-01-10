package main

import "fmt"

func main() {
	var uang int

	fmt.Scan(&uang)
	fmt.Print(uang / 10000)
	uang = (uang % 10000)
	fmt.Print(uang / 5000)
	uang = (uang % 5000)
	fmt.Print(uang / 1000)
}
