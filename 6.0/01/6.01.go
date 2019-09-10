package main

import "fmt"

func main() {
	x := 42
	fmt.Println(&x)
	foo(&x)
	fmt.Println(x)
}

func foo(x *int) {
	*x += 1
}
