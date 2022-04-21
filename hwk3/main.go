package main

import "fmt"

type Person struct {
	Firstname   string
	Lastname    string
	Phone       int
	Email       string
	MailAddress Address
	BillAddress Address
}

type Address struct {
	Address1 string
	Address2 string
	City     string
	State    string
	Zip      int
}

func (p *Person) Label() {

}

func alterPointer(AlterThisString *string, AlterThisPointer *string) {
	*AlterThisString = *AlterThisString + "abc"
	*AlterThisPointer = *AlterThisPointer + "abc"
}

func PassByPointer(name string, namebypointer *string) {
	fmt.Printf("Name=%v\n", name)
	fmt.Printf("Namebypointer=%v\n", namebypointer)
	alterPointer(&name, namebypointer)
	fmt.Printf("Name=%v\n", name)
	fmt.Printf("Namebypointer=%v\n", namebypointer)

}

// var x int
// var y bool
// var z string

// var m map [string]string
// var s Person
// var i interface{}

var intSlice []int

var boolSlice []bool

var stringSlice []string

func main() {
	// 	PrintIntSlice()
	// 	PrintBoolSlice()
	// 	PrintStringSlice()
	// 	PrintByteSlice()

	Evan := Person{
		Firstname: "Evan",
		Lastname:  "Kilpatrick",
		Phone:     7706886290,
		Email:     "ekilpatrick122495@gmail.com",
		MailAddress: Address{
			Address1: "804 Harbor Pointe Parkway",
			City:     "Roswell",
			State:    "GA",
			Zip:      30350,
		},
		BillAddress: Address{
			Address1: "804 Harbor Pointe Parkway",
			City:     "Roswell",
			State:    "GA",
			Zip:      30350,
		},
	}

	PassByPointer(Evan.Firstname, &Evan.Firstname)

	// fmt.Printf("%+v\n", Evan)

}

func PrintIntSlice() {
	intSlice = append(intSlice, 1, 2, 3, 4, 20, 3)

	for i := 0; i < len(intSlice); i++ {
		fmt.Println(intSlice[i])
	}

	for k, v := range intSlice {
		fmt.Printf("key: %v value: %v \n", k, v)
	}
}

func PrintBoolSlice() {
	//boolSlice = append(boolSlice, true, false, false, false, true, true)
	boolSlice = []bool{true, false, false, false, true, true}

	for i := 0; i < len(boolSlice); i++ {
		fmt.Println(boolSlice[i])
	}

	for k, v := range boolSlice {
		fmt.Printf("key: %v value: %v \n", k, v)
	}
}

func PrintStringSlice() {
	stringSlice = append(stringSlice, "true", "Evan", "Bob", "false", "VCS")

	for i := 0; i < len(stringSlice); i++ {
		fmt.Println(stringSlice[i])
	}

	for k, v := range stringSlice {
		fmt.Printf("key: %v value: %v \n", k, v)
	}
}

func PrintByteSlice() {
	ByteSlice := []byte("☺")
	ByteSlice = []byte("┐")

	for i := 0; i < len(ByteSlice); i++ {
		fmt.Println(ByteSlice[i])
	}

	for k, v := range ByteSlice {
		fmt.Printf("key: %v value: %v \n", k, v)
	}
}
