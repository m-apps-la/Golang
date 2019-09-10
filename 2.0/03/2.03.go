package main

import "fmt"

func main() {
	x := 1991
	for x < 2019 {
		fmt.Println(x)
		x += 1
	}
	fmt.Println(`Done`)
}
