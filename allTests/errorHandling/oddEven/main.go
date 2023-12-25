package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Pick a number")
		return
	}

	args, err := strconv.Atoi(os.Args[1])

	if err != nil  {
		fmt.Printf("%q is not a number\n", os.Args[1])
		return
	}

	if args % 8 == 0 {
		fmt.Printf("%d is an even number and it's divisible by 8", args)
	} else if args % 2 == 0 {
		fmt.Printf("%d is an even number", args)
	} else {
		fmt.Printf("%d is an odd number", args)
	}
}