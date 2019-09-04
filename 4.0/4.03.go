package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	vehicle
	fourWheels bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	x := truck{
		vehicle: vehicle{
			doors: 2,
			color: `Black`,
		},
		fourWheels: true,
	}
	y := sedan{
		vehicle: vehicle{
			doors: 4,
			color: `White`,
		},
		luxury: true,
	}

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(x.doors)
	fmt.Println(y.doors)
}
