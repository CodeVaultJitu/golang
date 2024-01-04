package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}

	// scale, err := strconv.ParseFloat(os.Args[1], 64)

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	switch s := os.Args[1]; s{
	case "massive":
		fmt.Printf(`%s's  richter scale is 10+`, s)
	case "great":
		fmt.Printf(`%s's  richter scale is 8 - 9.9`, s)
	case "major":
		fmt.Printf(`%s's  richter scale is 7 - 7.9`, s)
	case "strong":
		fmt.Printf(`%s's  richter scale is 6 - 6.9`, s)
	case "moderate":
		fmt.Printf(`%s's  richter scale is 5 - 5.9`, s)
	case "light":
		fmt.Printf(`%s's  richter scale is 4 - 4.9`, s)
	case "minor":
		fmt.Printf(`%s's  richter scale is 3 - 3.9`, s)
	case "very minor":
		fmt.Printf(`%s's  richter scale is 2 - 2.9`, s)
	case "micro":
		fmt.Printf(`%s's  richter scale is less than 2.0`, s)
	default:
		fmt.Printf(`%s's  richter scale is unknown`, s)
	}
}