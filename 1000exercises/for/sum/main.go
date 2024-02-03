package main

import "fmt"

func main() {

	// 2 ways
	
	var sum int
	for i := 1; i <= 10; i++ {
		for j := i + 1; j <= 10; j++ {
			sum = i + j
		}
	}
	fmt.Println(sum)

	var plus int
	for i := 1; i <= 10; i++ {
		plus += i
	}
	fmt.Println(plus)
}
