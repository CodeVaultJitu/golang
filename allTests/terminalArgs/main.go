package main

import (
	"fmt"
	"os"
)

func main() {
	// EXERCISE: Count the Arguments
	// count := len(os.Args)
	// fmt.Printf("There are %d names.\n", count)

	// EXERCISE: Print the Path
	// fmt.Println("Path is", os.Args[0])

	// EXERCISE: Print Your Name
	// name := os.Args[1]
	// fmt.Printf("Hi, %s\n How are you?", name)

	// EXERCISE: Greet More People
	// count := len(os.Args) - 1
	// if count < 3 {
	// 	fmt.Println("not 3 args")
	// 	return
	// } 
	// fmt.Printf("There are %d people.\n", count)
	// name := os.Args[1]
	// fmt.Printf("Hello great %s\n", name)
	// nam_e := os.Args[2]
	// fmt.Printf("Hello great %s\n", nam_e)
	// na_me := os.Args[3]
	// fmt.Printf("Hello great %s\n", na_me)

	// without using varibales
	fmt.Printf("There are %d people\n", len(os.Args)-1)
	fmt.Printf("Hello great %s\n", os.Args[1])
	fmt.Printf("Hello great %s\n", os.Args[2])
	fmt.Printf("Hello great %s\n", os.Args[3])
	fmt.Printf("Hello great %s\n", os.Args[4])
	fmt.Printf("Hello great %s\n", os.Args[5])
}
