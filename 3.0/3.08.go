package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	for k := range x {
		fmt.Println("This is the key: ", k)
		for nk, nv := range x[k] {
			fmt.Printf("\t %v likes %v: \t %v \n", k, nv, nk)
		}
	}
}
