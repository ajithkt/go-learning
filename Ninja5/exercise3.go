package main

import (
	"fmt"
)

type vehicle struct {
	doors string
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxuary bool
}

func main() {

	F150 := truck{
		vehicle: vehicle{
			doors: "four",
			color: "red",
		},
		fourWheel: true,
	}

	civic := sedan{
		vehicle: vehicle{
			doors: "four",
			color: "red",
		},
		luxuary: false,
	}

	fmt.Println(F150)
	fmt.Println(civic)
	fmt.Println(F150.doors)
	fmt.Println(civic.vehicle.color)

}
