package main

import "fmt"

// var x int
// var y bool
// var z string

// var m map [string]string
// var s Person
// var i interface{}

var intSlice []int

var boolSlice []bool

// var stringSlice []string

func main() {
	integer()
	boolean()
}

func integer() {
	intSlice = append(intSlice, 1, 2, 3, 4, 20, 3)

	for i := 0; i < len(intSlice); i++ {
		fmt.Println(intSlice[i])
	}

	for k, v := range intSlice {
		fmt.Printf("key: %v value: %v \n", k, v)
	}
}

func boolean() {
	boolSlice = append(boolSlice, true, false, false, false, true, true)

	for i := 0; i < len(boolSlice); i++ {
		fmt.Println(boolSlice[i])
	}

	for k, v := range boolSlice {
		fmt.Printf("key: %v value: %v \n", k, v)
	}
}
