package main

import (
	"fmt"
	"os"
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

// func isVowel(args string) string {

// 	//indexany, containsany
// 	detect := strings.IndexAny(args, "aeiou")

// 	if detect != -1 {
// 		return fmt.Sprintf("%q is a vowel\n", args[0])
// 	} else if args[0] == 'y' || args[0] == 'w' {
// 		return fmt.Sprintf("%q is sometimes a vowel, sometimes not\n", args[0])
// 	} else {
// 		return fmt.Sprintf("%q is a consonant\n", args[0])
// 	}
// }

// func main () {
// args := os.Args[1:]

// 	if len(args) != 1 || len(args[0]) != 1 {
// 		 fmt.Printf("Give me a letter\n")
// 	} else {
// 		result := isVowel(args[0])
// 		fmt.Println(result)
// 	}
// }

	func passwordProtect(args []string) string {
		if args[0] == "jack" && args[1] == "1888" {
			return fmt.Sprintf("Access granted for %q", args[0])
		} else if args[0] == "jack" && args[1] != "1888" {
			return fmt.Sprintf("Invalid password for %q", args[0])
		} else {
			return fmt.Sprintf("Access denied for %q", args[0])
		}
	}

func main () {
		args := os.Args[1:]

		if len(args) != 2 {
			fmt.Println("Usage: [username] [password]")
			return
		} else {
			result := passwordProtect(args)
			fmt.Println(result)
		}

}