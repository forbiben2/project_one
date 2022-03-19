package main

import (
	"fmt"
)

func main() {

	for day := 0; day < 13; day++ {
		if day == 0 {
			fmt.Printf("The 12 Days of Christmas\n\n")
			day = day + 1
		}
		if day == 1 {
			fmt.Println("On the", day, "st day of christmas\nMy true love gave to me")
		} else if day == 2 {
			fmt.Println("On the", day, "nd day of christmas\nMy true love gave to me")
		} else if day == 3 {
			fmt.Println("On the", day, "rd day of christmas\nMy true love gave to me")
		} else {
			fmt.Println("On the", day, "th day of christmas\nMy true love gave to me")
		}

		switch day {
		case 12:
			fmt.Println("Twelve drummers drumming,")
			fallthrough
		case 11:
			fmt.Println("Eleven pipers pipping,")
			fallthrough
		case 10:
			fmt.Println("Ten lords a leaping,")
			fallthrough
		case 9:
			fmt.Println("Nine ladies dancing,")
			fallthrough
		case 8:
			fmt.Println("Eight maids a milking,")
			fallthrough
		case 7:
			fmt.Println("Seven swans a swimming,")
			fallthrough
		case 6:
			fmt.Println("Six geese-a-laying,")
			fallthrough
		case 5:
			fmt.Println("Five golden rings,")
			fallthrough
		case 4:
			fmt.Println("Four calling bird,")
			fallthrough
		case 3:
			fmt.Println("Three french hens,")
			fallthrough
		case 2:
			fmt.Println("Two turtle doves and,")
			fallthrough
		case 1:
			fmt.Println("A partridge in a pear tree \n \n ")
			if day == 12 {
				break
			}
		}

	}
}
