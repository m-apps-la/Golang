package main

import "fmt"

func main() {
	x := 0
	for x < 10 {
		if x == 10 {
			break
		}
		fmt.Println(x)
		x += 1
	}
	fmt.Println(`Done`)
}
