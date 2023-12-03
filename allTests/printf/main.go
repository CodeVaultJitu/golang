package main

import (
	"fmt"
	"os"
)

func main() {
	// fmt.Printf("I'm 30 years old")

	// name := "Jitu"
	// lastName := "Hazarika"
	// format := "My name is %v and my lastname is %v."
	// fmt.Printf(format, name, lastName)

	// tf := false
	// fmt.Printf("These are %t claims", tf)

	// temp := 25.1
	// fmt.Printf("Temperature is %0.1f degrees", temp)

	// hello := "hello world"
	// fmt.Printf("%q\n", hello)

	// value := 3
	// fmt.Printf("Type of %d is %T", value, value)

	// value := 3.14
	// fmt.Printf("Type of %0.2f is %T", value, value)	

	// value := "hello"
	// fmt.Printf("Type of %s is %T", value, value)	

	// value := true
	// fmt.Printf("Type of %t is %T", value, value)	

	value := "Your name is %s and your lastname is %s"
	fName, lName := os.Args[1], os.Args[2]
	fmt.Printf(value, fName, lName)

}