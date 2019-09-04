package main

import "fmt"

func main() {
	x := struct {
		name      string
		age       int
		friends   map[string]int
		favDrinks []string
	}{
		name: `Mike`,
		age:  27,
		friends: map[string]int{
			`John`: 12345,
			`Jake`: 23456,
		},
		favDrinks: []string{`Gin`, `Scotch`},
	}

	fmt.Println(x)
	fmt.Println(x.name)
	fmt.Println(x.age)
	fmt.Println(x.friends)
	fmt.Println(x.friends[`John`])
	fmt.Println(x.favDrinks)
	fmt.Println(x.favDrinks[1])

	for k, v := range x.friends {
		fmt.Println(k, v)
	}
	for i, v := range x.favDrinks {
		fmt.Println(i, v)
	}
}
