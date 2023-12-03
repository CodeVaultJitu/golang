package main

import "fmt"

type contactInfo struct {
	email   string
	pincode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// alex := person{firstName: "Alex", lastName: "Alek"}
	// alex.firstName = "anderson"
	// fmt.Println(alex)
	Troy := person{
		firstName: "Troy",
		lastName:  "Alek",
		contact: contactInfo{
			email:   "troy@gmail.com",
			pincode: 785684,
		},
	}

// Troy.update()
Troy.update("Jax")
Troy.print()
}

func (p *person) update(newFirstName string){
	p.firstName = newFirstName
}

// func (p *person) update(){
// 	p.firstName= "Jax"
// }

func (p person) print(){
	fmt.Println(p)
}
