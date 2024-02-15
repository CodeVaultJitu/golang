package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const maxTurns = 5

var n int

func main() {
	// rand.NewSource(time.Now().UnixNano())

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("provide a number")
		return
	}

	guess, err := strconv.Atoi(args[0])
	if err != nil || guess < 0 {
		fmt.Println("Not a number")
		return
	}

	for turn := 1; turn < maxTurns; turn++ {
		n = rand.Intn(guess + 1)

		// to not use nested if loop
		// if n != guess {
		// 	continue
		// }

		// if turn == 1 {
		// 	fmt.Println("This is your first turn and You WON !!!")

		// } else {
		// 	fmt.Println("You win")
		// }
		// return
		if n == guess {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("Winner")
			case 1:
				fmt.Println("You're awesome")
			case 2:
				fmt.Println("Perfect")
			}
			return
		}
	}
	// out of for loop to run upto max turns

	switch rand.Intn(2) {
	case 0:
		fmt.Println("Loser")
	case 1:
		fmt.Println("Just bad luck")
	}
}
