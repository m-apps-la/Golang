package main

import "fmt"

func main() {
	for i := 10; i <= 100; i += 1 {
		fmt.Printf("When %v is mod4 the remainder is %v\n", i, i%4)
	}
}
