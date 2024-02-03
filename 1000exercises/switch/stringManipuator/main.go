package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println("[command] [string]")
        fmt.Println("Available commands: lower, upper and title")
        return
    }

    s := os.Args[2]

    switch o := os.Args[1]; o {
    case "lower":
        fmt.Println(strings.ToLower(s))
    case "upper":
        fmt.Println(strings.ToUpper(s))
    case "title":
			//strings.toTitle is depreciated
        fmt.Println(cases.Title(language.English).String(s)) // Use language.English for title casing
    default:
        fmt.Printf("Unknown command: %q", o)
    }
}
