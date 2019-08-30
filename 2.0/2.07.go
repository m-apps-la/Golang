package main

import "fmt"

func main() {
	x := "Tnaya"

	if x == "Mike" {
		fmt.Println(`x = "Mike"`)
	} else if x == "Tnaya" {
		fmt.Println(`x = "Tnaya"`)
	} else {
		fmt.Println(`x = "IDK"`)
	}
}
