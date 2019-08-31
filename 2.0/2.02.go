package main

import "fmt"

func main() {
	for i := 65; i < 91; i += 1 {
		fmt.Println(i)
		for j := 0; j < 3; j += 1 {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
