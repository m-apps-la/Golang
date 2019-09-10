package main

import "fmt"

func main() {
	switch "favSport" {
	case "Mike":
		fmt.Println(`Should not Print`)
	case "favSport":
		fmt.Println(`Racing`)
	default:
		fmt.Println(`Default`)
	}
}
