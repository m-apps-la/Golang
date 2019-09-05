package main

import "fmt"

func main() {
	defer foo()
	bar()
}

func foo() {
	fmt.Println(`Winning`)
}

func bar() {
	fmt.Println(`Mike is`)
}
