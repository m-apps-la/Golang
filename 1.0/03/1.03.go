package main

import "fmt"

const (
	a     = 2
	b int = 3
)

func main() {
	fmt.Printf("%T\t%T\t", a, b)
	fmt.Print(a, b)
}
