package main

import "fmt"

type user struct {
	first string
	last  string
	age   int
}

func (s user) speak() {
	fmt.Printf("%v %v is %v years old!", s.first, s.last, s.age)
}

func main() {
	x := user{
		first: `Mike`,
		last:  `Mulderink`,
		age:   27,
	}
	x.speak()
}
