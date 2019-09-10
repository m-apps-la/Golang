package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	for i := 0; i < 10; i += 1 {
		go func() {
			for i := 0; i < 10; i += 1 {
				c <- i
			}
		}()
	}

	for j := 0; j < 100; j += 1 {
		fmt.Println(j, <-c)
	}

	close(c)
}
