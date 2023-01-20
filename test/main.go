package main

import "fmt"

type person struct {
	name string
}

func main() {
	p := person{"Richard"}
	rename(p)
	fmt.Println(p)
}
func rename(p person) {
	p.name = "test"
}
