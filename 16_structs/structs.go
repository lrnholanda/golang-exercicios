package main

import "fmt"

type Address struct {
	city   string
	street string
	postal string
}

type Person struct {
	name    string
	address Address
}

// Relying on default naming
type Employee struct {
	Address
	company string
}

var address Address
var person Person

func main() {
	// Embedding a struct
	person := Person{
		name: "chris",
		address: Address{
			city: "Stockholm",
		},
	}

	fmt.Println(person)

	employee := Employee{
		Address: Address{
			city: "LA",
		},
		company: "Microsoft",
	}

	fmt.Println(employee)

	address.city = "London"
	address.street = "Buckingham palace"
	address.postal = "SW1"
	fmt.Println(address.GetAdress())

	person.name = "Lorena"
	person.address.city = "Caxias"
	fmt.Println(person.GetPerson())
}

func (a Address) GetAdress() string {
	return fmt.Sprintf("City: %s, Street: %s, Postal address: %s", a.city, a.street, a.postal)
}

func (p Person) GetPerson() string {
	return fmt.Sprintf("Name: %s, City:  %s", p.name, p.address.city)
}
