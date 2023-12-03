package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":    "#ffffff",
		"purple": "#4bf",
		"white": "#df3465",
	}

	// var colors map[string]string
		// colors := make(map[string]string)
		// colors["white"] = "#ght456se"
		// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c{
		fmt.Println("color is ", color, "and hex is ", hex)
	}
}