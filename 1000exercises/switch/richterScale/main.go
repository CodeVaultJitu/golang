package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}

	scale, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil {
		fmt.Println(err)
		return
	}

	switch s := scale; {
	case s >= 10.0:
		fmt.Printf("%.2f is massive", s)
	case s<10.0 && s>=8.0:
		fmt.Printf("%.2f is great", s)
	case s<8.0 && s>=7.0:
		fmt.Printf("%.2f is major", s)
	case s<7.0 && s>=6.0:
		fmt.Printf("%.2f is strong", s)
	case s<6.0 && s>=5.0:
		fmt.Printf("%.2f is moderate", s)
	case s<5.0 && s>=4.0:
		fmt.Printf("%.2f is light", s)
	case s<4.0 && s>=3.0:
		fmt.Printf("%.2f is minor", s)
	case s<3.0 && s>=2.0:
		fmt.Printf("%.2f is very minor", s)
	default:
		fmt.Printf("%.2f is micro", s)
	}
}