package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println(`Please enter 2 numbers
The first number should be less than the 2nd number`)
		return
	}

	min, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Please enter 2 numbers")
		return
	}

	max, errr := strconv.Atoi(os.Args[2])
	if errr != nil {
		fmt.Println("Please enter 2 numbers")
		return
	}

	if min < max {
		var plus int
		for i := min; i <= max; i++ {
			if i == max {
				fmt.Printf("%d = ", i)
				//printing = after the last number
			} else {
				fmt.Printf("%d + ", i)
				// will print + after every number until
			}
			plus += i
		}
		fmt.Println(plus)
	} else {
		fmt.Println("The first number should be less than the 2nd number")
		return
	}

}
