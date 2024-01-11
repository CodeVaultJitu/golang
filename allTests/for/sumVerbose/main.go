package main

import "fmt"

func main() {

	var plus int
	for i := 1; i <= 10; i++ {
		if i == 10 {
			fmt.Printf("%d = ", i)
		} else {
			fmt.Printf("%d + ", i)
		}
		plus += i
	}
	fmt.Println(plus)
	
}