package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) < 2 {
		if len(args) == 1 {
			_, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Size is missing")
				fmt.Println("Usage: [op=*/+-] [size]")
			} else {
				fmt.Println("Operator is missing")
				fmt.Println("Usage: [op=*/+-] [size]")
			}
		} else {
			fmt.Println("Usage: [op=*/+-] [size]")
		}
		return
	}

	operator := args[0]
	size := args[1]

	if strings.IndexAny(operator, "=*/+-%") == -1 {
		fmt.Println(`Invalid operator.
Valid ops one of: */+-%`)
		return
	}

	s, err := strconv.Atoi(size)
	if err != nil {
		fmt.Println("Invalid size. Size should be an integer.")
		os.Exit(1)
	}

	fmt.Printf("%5s", operator)
	for i := 0; i <= s; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= s; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= s; j++ {
			var res int

			switch operator {
			case "*":
				res = i * j
			case "%":
				if j != 0 {
					res = i % j
				}
			case "/":
				if j != 0 {
					res = i / j
				}
			case "+":
				res = i + j
			case "-":
				res = i - j

			}

			fmt.Printf("%5d", res)
		}
		fmt.Println()
	}
	}
