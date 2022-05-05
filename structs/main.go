package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"}

	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println((alex.lastName))

	// var alex person
	// alex.firstName, alex.lastName = "Alex", "Anderson"
	// fmt.Println(alex.firstName)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Palmetto",
		contactInfo: contactInfo{
			email: "Jimpalmetto@gmail.com",
			zip:   123456,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	fmt.Printf("%+v", jim)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	*&pointerToPerson.firstName = newFirstName
}
