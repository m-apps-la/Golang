package main

import "fmt"

func main() {
	x := 1991
	for {
		if x > 2019 {
			break
		}
		fmt.Println(x)
		x += 1
	}
	fmt.Print(`Done`)
}
