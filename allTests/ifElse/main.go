package main

import (
	"fmt"
	"os"
)

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

	// isSphere, radius := true, 200
	// var big bool

	// big = isSphere && radius >= 200

	// if big {
	// 	fmt.Println("It's a big sphere.")
	// } else {
	// 	fmt.Println("I don't know.")
	// }

		args := os.Args[1:]
		count := len(args)

		if count < 1 {
			fmt.Println("Give me args")
		}else if count == 1 {
			//q is for quoted outputs, without args[] it will print with sq bracks to represent
			//the slice of array in which it is stored
			fmt.Printf("There is one: %q", args[0])
		} else if count == 2 {
				fmt.Printf("There are 2: %q", args[0]+" "+args[1])
		} else {
			fmt.Printf("There are %d arguments", count)
		}
}