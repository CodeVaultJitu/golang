package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	// 	msg := `

	// 	The weather looks good.
	// I should go and play.

	// 	`
	// 	fmt.Println(strings.TrimSpace(msg))

	name := "inanç           "
	name = strings.TrimRight(name, " ")
	fmt.Printf("%s\n", name)
	fmt.Println(utf8.RuneCountInString(name))
}