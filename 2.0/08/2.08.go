package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println(`Should not Print`)
	case true:
		fmt.Println(`Sould Print`)
	default:
		fmt.Println(`Default - Should not Print`)
	}
}
