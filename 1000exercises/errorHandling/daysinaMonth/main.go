package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//checking if there's arguments in the cli or not
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	//taking current time and year
	year := time.Now().Year()
	//leap year check
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	if m := strings.ToLower(os.Args[1]);  m == "january" ||
	  m ==  "march" ||
	  m ==  "may" ||
	  m ==  "july" ||
	  m ==  "august" ||
	  m ==  "october" ||
	  m ==  "december" {
		fmt.Printf("%q has 31 days.", os.Args[1])
	} else if m == "february" && leap {
		fmt.Printf("%q has 29 days", os.Args[1])
	} else if m == "february" {
		fmt.Printf("%q has 28 days", os.Args[1])
	} else if m ==  "april" ||
	m ==  "june" ||
	m ==  "september" ||
	m ==  "november" {
		fmt.Printf("%q has 30 days", os.Args[1])
	} else {
		fmt.Printf("%q is not a month.", os.Args[1])
	}
}