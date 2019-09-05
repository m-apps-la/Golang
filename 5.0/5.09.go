package main

import "fmt"

func main() {
	foo(cb, 42)
}

func foo(x func() string, y int) {
	fmt.Print(x())
	fmt.Printf(" %v times!", y)
}

func cb() string {
	return fmt.Sprint(`Mike has Won`)
}
