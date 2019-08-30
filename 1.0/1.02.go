package main

import "fmt"

func main() {
	a := (2 == 2)
	b := (3 <= 2)
	c := (3 >= 2)
	d := (2 != 2)
	e := (3 < 2)
	f := (3 > 2)
	fmt.Println(a, b, c, d, e, f)
}
