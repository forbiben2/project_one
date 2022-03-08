package main

import (
	"fmt"
)

func main() {
	x := 10
	// for i := 1; x < 20; i++ {
	// 	fmt.Println(x)
	// 	x += i
	// }

	if x < 10 {
		fmt.Printf("x is less than 10\n")
	} else if x > 10 {
		fmt.Printf("x is greater than 10\n")
	} else {
		fmt.Printf("x is 10\n")
	}
}
