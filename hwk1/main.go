 package main

 import ("fmt"
"strings"
)

func main() {
    nameAsString()
	lastNameAsString()
	ageAsInt()
	ageDivide3Int()
	ageDivide3Remainder()
	ageDivide3Float()
	address()
	phone()
	displayPersonalStats(nameAsString(), lastNameAsString(), ageAsInt())
	printStats(ageAsInt(), nameAsString(), lastNameAsString(), phone(),address())
	envelopeCover(nameAsString(), lastNameAsString(),address())
    return
}

func nameAsString() (x string) {
	//x = "Evan"
	x = strings.ToLower("Evan")
    return x
}

func lastNameAsString() (y string) {
	y = strings.ToLower("Kilpatrick")
    return y
}

func ageAsInt() (z int) {
	return 26
}

func ageDivide3Remainder() {
	fmt.Printf("remainder: %v \n", 26%3)
}

func ageDivide3Int() {
	fmt.Printf("Age/3: %v ", 26/3)
}

func ageDivide3Float() {
	var x,y int =26,3
	fmt.Printf("Age/3 in decimal form: %v \n", float64(x)/float64(y))
}

func address() [5]string{
	var address [5]string
		address[0]="217 barrington oaks ridge"
		address[1]="unit 217"
		address[2]="roswell"
		address[3]="georgia"
		address[4]="30075"
	return address
}

func phone() string{
	return "404-123-4567"
}

func displayPersonalStats(firstname string, lastname string, age int) {
	firstname=strings.Title(firstname)
	lastname=strings.Title(lastname)
	fmt.Printf("Name: %v %v\n", firstname, lastname)
	fmt.Printf("Age: %v\n", age)

}
func printStats(age int, firstname string, lastname string, phone string,address [5]string){
	firstname=strings.Title(firstname)
	lastname=strings.Title(lastname)
	w:=strings.Title(address[0])
	a:=strings.Title(address[1])
	x:=strings.Title(address[2])
	y:=strings.Title(address[3])
	z:=strings.Title(address[4])
	fmt.Printf("Name:%v %v \nAge:%v \nAddress:%v %v\n%v, %v, %v \nPhone:%v \n",firstname, lastname, age, w, x, y, z, a, phone)
}

func envelopeCover(firstname,lastname string,address [5]string){
	firstname=strings.Title(firstname)
	lastname=strings.Title(lastname)
	w:=strings.Title(address[0])
	a:=strings.Title(address[1])
	x:=strings.Title(address[2])
	y:=strings.Title(address[3])
	z:=strings.Title(address[4])
	fmt.Printf("%v %v\n%v\n%v\n%v,%v,%v \n",firstname, lastname, w, x, y, z, a)
}