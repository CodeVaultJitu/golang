package main

import (
	"fmt"
	"time"
)

func main() {
	switch 	t := time.Now().Hour();	{
	case t > 19:
		fmt.Printf("good night, the time is %d \n", t)
	case t > 12:
		fmt.Printf("good afternoon, the time is %d \n", t)		
	default:
		fmt.Printf("good morning, the time is %d \n", t)
	}
}