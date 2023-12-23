package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	if  len(os.Args) < 2 {
		fmt.Println("Requires age")
		return
	}
	
	ageStr := os.Args[1]

	age, err := strconv.Atoi(ageStr)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

 if age <= 0 {
		fmt.Printf(`Wrong age: "%d" `, age)
	} else if age > 17 {
		fmt.Printf("R-Rated")
	} else if age >= 13 && age <= 17 {
		fmt.Println("PG-13")
	} else {
		fmt.Println("PG-Rated")
	}
}