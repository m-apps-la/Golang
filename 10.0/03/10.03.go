package main

import "fmt"

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the Error, %v", ce.info)
}

func main() {
	c1 := customErr{
		info: `Mike needs more Single Malt Scotch`,
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println("Foo ran -", e)
	fmt.Println("Foo ran -", e.(customErr).info)
}
