package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Println(`Please enter 2 numbers
The first number should be less than the 2nd number`)
		return
	}

	min, err := strconv.Atoi(os.Args[1])
	max, errr := strconv.Atoi(os.Args[2])

	if err != nil || errr != nil {
		fmt.Println(`Please enter 2 numbers
The first number should be less than the 2nd number`)
		return
	} else if min > max {
		fmt.Println("The first number should be less than the 2nd number")
		return
	}

	var plus int
	i := min

	for {

		if i > max {
			break
		} else if i%2 != 0 {
			i++
			continue
		}

		plus += i

		if i == max && i%2 == 0 {
			fmt.Printf("%d = ", i)
		} else if i < max && i == max-1 {
			fmt.Printf("%d = ", i)
		} else if i < max {
			fmt.Printf("%d + ", i)
		}
		i++
	}
	fmt.Printf("%d", plus)

}
