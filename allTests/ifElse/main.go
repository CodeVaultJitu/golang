package main

import (
	"fmt"
	"os"
	"strings"
)

// func main() {
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

// args := os.Args[1:]
// count := len(args)

// if count < 1 {
// 	fmt.Println("Give me args")
// }else if count == 1 {
// 	//q is for quoted outputs, without args[] it will print with sq bracks to represent
// 	//the slice of array in which it is stored
// 	fmt.Printf("There is one: %q", args[0])
// } else if count == 2 {
// 		fmt.Printf("There are 2: %q", args[0]+" "+args[1])
// } else {
// 	fmt.Printf("There are %d arguments", count)
// }

// }

func isVowel(args string) string {

	detect := strings.ContainsAny(args, "aeiou")

	if detect {
		return fmt.Sprintf("%q is a vowel\n", args[0])
	} else if args[0] == 'y' || args[0] == 'w' {
		return fmt.Sprintf("%q is sometimes a vowel, sometimes not\n", args[0])
	} else {
		return fmt.Sprintf("%q is a consonant\n", args[0])
	}
}

func main () {
	args := os.Args[1:]

		if len(args) != 1 || len(args[0]) != 1 {
			 fmt.Printf("Give me a letter\n")
		} else {
			result := isVowel(args[0])
			fmt.Println(result)
		}
}