package main

import "fmt"

func main() {
	foo()

	c := make(chan int)

	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
}

func foo() {
	c := make(chan int, 1)

	c <- 43

	fmt.Println(<-c)
}
