package main

import "fmt"

func main() {
	color, color2 := "red", "blue"
	color, color2 = "orange", "green"
	fmt.Println(color, color2)

	red, blue := "red", "blue"
	red, blue = blue, red
	fmt.Println(red, blue)
}