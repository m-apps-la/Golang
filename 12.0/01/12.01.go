package main

import (
	"fmt"

	"github.com/Golang/12.0/01/cat"
)

type serval struct {
	name string
	age  int
}

func main() {
	fido := sevval{
		name: "Fido",
		age:  cat.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(cat.YearsTwo(20))
}
