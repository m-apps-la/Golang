package main

import "fmt"

type person struct {
	first string
}

func changeMe(p *person) {
	p.first = `Tnaya`
	(*p).first = `Tnaya W`
}

func main() {
	p := person{
		first: `Mike`,
	}
	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)
}
