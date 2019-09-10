package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%T\n%T\n%T\n", x, y, z)
	a := 0
	if a == 0 {
		a = 2
	}
	if a == 0 {
		a = 2
	}
	if a <= 2 {
		a = 5
	}
	if a >= 4 {
		a = 1
	}
	if a != 0 {
		a = 4
	}
	if a > 3 {
		a = 6
	}
	if a < 7 {
		a = 5
	}
	fmt.Println(a)
}
