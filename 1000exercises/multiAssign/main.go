package main

import "fmt"

func main() {
	var (
		// a, b, c string
		planet string
		isTrue bool
		temp float64
	)

	// a = "air is good on mars"
	// b = "it's true"
	// c = "it is 19.5 degrees"

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	planet = "air is good on mars"
	fmt.Println(planet)

	planet = "it's"
	isTrue = true
	fmt.Println(planet, isTrue)

	temp = 19.5
	output := fmt.Sprintf("it is %.1f degrees", temp) 
	fmt.Println(output)

}
