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

	m := map[string]person{
		x.first: x,
		y.first: y,
	}

	fmt.Println()

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for k, nv := range v.fav {
			fmt.Printf("%v %v \n", k, nv)
		}
		fmt.Println()
	}
}
