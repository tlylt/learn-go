package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	// contactInfo
}

// func main() {
// 	alex := person{"Alex", "Anderson"}
// }

// func main() {
// 	alex := person{firstName: "Alex", lastName: "Anderson"}
// 	fmt.Println(alex)
// }

// func main() {
// 	var alex person
// 	alex.firstName = "Alex"
// 	alex.lastName = "Anderson"
// 	fmt.Println(alex)
// 	fmt.Printf("%+v", alex)
// }

// func main() {
// 	jim := person{
// 		firstName: "Jim",
// 		lastName:  "Party",
// 		contact: contactInfo{
// 			email:   "jim@gmail.com",
// 			zipCode: 94000,
// 		},
// 	}
// 	jimPointer := &jim
// 	jimPointer.updateName("Jimmy")
// 	jim.print()
// }

// func main() {
// 	jim := person{
// 		firstName: "Jim",
// 		lastName:  "Party",
// 		contact: contactInfo{
// 			email:   "jim@gmail.com",
// 			zipCode: 94000,
// 		},
// 	}
	
// 	jim.updateName("Jimmy")
// 	jim.print()
// }

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func main() {
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	// var colors map[string]string

	colors := make(map[string]string)

	colors["red"] = "#ff0000"

	fmt.Println(colors)
}