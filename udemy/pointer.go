package main

import "fmt"

type ContactInfo struct {
	email_address string
	zip_code      int
}

type Person struct {
	firstName string
	lastName  string
	ContactInfo
}

func (person Person) print() {
	fmt.Printf("%v\n", person)
}

func (person *Person) updateName(first_name string) {
	(*person).firstName = first_name
}

func main() {
	person := Person{
		firstName: "Manoj",
		lastName:  "Kumar",
		ContactInfo: ContactInfo{
			email_address: "manutvm07@gmail.com",
			zip_code:      695573,
		},
	}

	person.print()
	(&person).updateName("Manu")
	person.print()
	person.updateName("Manikkkuttan")
	person.print()
}
