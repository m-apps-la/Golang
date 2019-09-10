package main

import "fmt"

func main() {
	x := []int{3, 6, 9, 38, 56, 5}
	fmt.Println(foo(x...))
	fmt.Println(bar(x))
}

func foo(x ...int) int {
	var j int
	for _, v := range x {
		j += v
	}
	return j
}

func bar(x []int) int {
	j := 0
	for _, v := range x {
		j += v
	}
	return j
}
