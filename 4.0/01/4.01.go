package main

import "fmt"

type person struct {
	first string
	last  string
	fav   []string
}

func main() {
	x := person{
		first: "Mike",
		last:  "Mulderink",
		fav:   []string{"Winning", "Happy"},
	}
	y := person{
		first: "Tnaya",
		last:  "Witmer",
		fav:   []string{"Mike", "Charlie", "Peter"},
	}

	fmt.Println(x.first)
	fmt.Println(x.last)
	fmt.Println()
	for _, v := range x.fav {
		fmt.Println(v)
	}

	fmt.Println()
	fmt.Println()

	fmt.Println(y.first)
	fmt.Println(y.last)
	fmt.Println()
	for _, v := range y.fav {
		fmt.Printf("Tnaya likes %v in this order \n", v)
	}
}
