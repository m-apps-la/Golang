package main

import "fmt"

func main() {
	x := []int{3, 56, 0, -5, 58, 3, 95, -289, 6, 78}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", x)
}
