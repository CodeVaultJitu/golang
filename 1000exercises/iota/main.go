package main

import "fmt"

// func main() {
// 	const(
// 		Nov = 11 - iota
// 		Oct
// 		Sep
// 	)

// 	fmt.Println(Sep, Oct, Nov)
// }

// func main() {
// 	const _ = iota
// 	const(
// 		Jan = iota + 1
// 		Feb
// 		Mar
// 	)
// 	fmt.Println(Jan, Feb, Mar)
// }

func main() {
	const(
		Winter = 12 - (3*iota)
		Fall
		Summer
		Spring
	)
	fmt.Println(Winter, Spring, Summer, Fall)
}