 package main

 import ("fmt"
"strings"
)

func main() {
    nameAsString()
	lastNameAsString()
	ageAsInt()
	//ageDivide3Int()
	//ageDivide3Remainder()
	//ageDivide3Float()
	address()
	phone()
	//displayPersonalStats(nameAsString(), lastNameAsString(), ageAsInt())
	printStats(ageAsInt(), nameAsString(), lastNameAsString(), phone())
	envelopeCover(nameAsString(), lastNameAsString())
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

func address() (string,string,string,string,string){
	w :="217 barrington oaks ridge"
	x :="unit 217"
	y :="roswell"
	z :="georgia"
	a :="30075"
	return w, x, y, z, a
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
func printStats(age int, firstname string, lastname string, phone string){
	w,x,y,z,a :=address()
	firstname=strings.Title(firstname)
	lastname=strings.Title(lastname)
	w=strings.Title(w)
	a=strings.Title(a)
	x=strings.Title(x)
	y=strings.Title(y)
	z=strings.Title(z)
	fmt.Printf("Name:%v %v \nAge:%v \nAddress:%v %v\n%v, %v, %v \nPhone:%v \n",firstname, lastname, age, w, x, y, z, a, phone)

}
func envelopeCover(firstname string, lastname string){
	w,x,y,z,a :=address()
	firstname=strings.Title(firstname)
	lastname=strings.Title(lastname)
	w=strings.Title(w)
	a=strings.Title(a)
	x=strings.Title(x)
	y=strings.Title(y)
	z=strings.Title(z)
	fmt.Printf("%v %v\n%v\n%v\n%v,%v,%v \n",firstname, lastname, w, x, y, z, a)
}
