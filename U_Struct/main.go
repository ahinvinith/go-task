package main

import "fmt"

type contactInfo struct {
	email   string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {

	// example for declaring a embeded struct

	ahin := person{
		firstName: "Ahin",
		lastName:  "Vinith",
		contact:   contactInfo{email: "ahin@gmail.com", pinCode: 629177},
	}

	
    // shortcut with GO pointer
	ahin.updateName("Prem")
	ahin.printPerson()

	// pointer example
	// ahinPointer := &ahin
	// ahinPointer.updateName("Prem")
	// ahin.printPerson()

	// fmt.Println(ahin)
	// fmt.Printf("My personal info is %+v \n", ahin)
	// fmt.Printf("%v", ahin)

	// 3rd way of declaring a struct
	// var ahin person
	// ahin.firstName = "Ahin"
	// ahin.lastName = "Vinith"

	//different types of declaring a struct
	// fmt.Println(ahin)
	// fmt.Printf(" %+v\n", ahin)
	// fmt.Printf(" %v\n", ahin)

	// ahin := person{firstName: "Ahin", lastName: "Vinith"}
	// fmt.Println("My name is ", ahin)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) printPerson() {
	fmt.Println(p)
	fmt.Printf("My personal info is %+v \n", p)
	fmt.Printf("%v\n", p)
}
