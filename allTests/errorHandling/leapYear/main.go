package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a year number")
		return
	}

	n, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Printf("%q is not a valid year.", os.Args[1])
		return
	}

	if n%4 == 0 && n%100 != 0 || n%400 == 0 {
		fmt.Printf("%d is a leap year.", n)
	} else {
		fmt.Printf("%d is not a leap year.", n)
	}

}