package main

import "fmt"

func main() {
	x := 0
	for {
		if x == 13 {
			break
		}
		fmt.Println(x)
		x += 1
	}
	fmt.Print(`Done`)
}
