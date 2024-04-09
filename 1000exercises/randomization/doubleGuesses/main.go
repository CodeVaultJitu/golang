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
	if len(args) < 2 {
		fmt.Println("provide a number")
		return
	}

	guess, err := strconv.Atoi(args[0])
	guess1, err1 := strconv.Atoi(args[1])
	if err != nil || err1 != nil || guess < 0 || guess1 < 0 {
		fmt.Println("Not a number")
		return
	}

	for turn := 1; turn < maxTurns; turn++ {
		n = rand.Intn(guess1 + 1)

		if n == guess || n == guess1 {
			switch rand.Intn(3) {
			case 0:
				fmt.Printf("Winner, number %d", n)
			case 1:
				fmt.Printf("You're awesome, number %d", n)
			case 2:
				fmt.Printf("Perfect, number %d", n)
			}
			return
		}
	}
	// out of for loop to run upto max turns

	switch rand.Intn(2) {
	case 0:
		fmt.Printf("Loser, number %d", n)
	case 1:
		fmt.Printf("Just bad luck, number %d", n)
	}
}
