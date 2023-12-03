package main

import "fmt"

func main() {
	// age := 10
	// if age > 60 {
	// 	fmt.Println("gettig older")
	// } else if age > 30 {
	// 	fmt.Println("Getting wiser")
	// } else if age > 20 {
	// 	fmt.Println("Adulthood")
	// } else if age > 10 {
	// 	fmt.Println("Young blood")
	// } else {
	// 	fmt.Println("Booting up")
	// }

	isSphere, radius := true, 200
	var big bool

	big = isSphere && radius >= 200

	if big {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}


}