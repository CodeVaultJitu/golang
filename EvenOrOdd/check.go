package main

import (
	"fmt"
)

func print(check []int) {
	for _, value := range check {
		if value % 2 == 0 {
			fmt.Println(value, "Even")
		} else {
			fmt.Println(value, "Odd")
		}
	}
}