package main

import "fmt"

func main() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)
}

// EXPECTED OUTPUT
//  10.5

//func main() {
 // a, b := 10, 5.5
 //a = int(b)
  //fmt.Println(float64(a) + b)
 // }

// EXPECTED OUTPUT
//  5.5
// ---------------------------------------------------------

//func main() {
 // fmt.Println(float64(5.5))
 //}
 
 // EXPECTED OUTPUT
//  9.5
// ---------------------------------------------------------

//func main() {
  // age := 2
  // fmt.Println(float64(7.5) + float64(age))
//}

// EXPECTED OUTPUT
//  1127
// ---------------------------------------------------------

//func main() {
  
  //min := int8(127)
 // max := int16(1000)

 // fmt.Println(max + int16(min))
//}