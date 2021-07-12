package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Sex       string
	Country   string
}

type Pet struct {
	Species string
	Name    string
	Age     int
	Color   string
}

type Connection interface {
	Attach(person *Person) string
}

func (pet Pet) Attach(person *Person) string {
	return pet.Name + " is " + person.FirstName + " " + person.LastName + "'s pet."
}

func main() {
	august := Pet{
		Species: "Dog",
		Name:    "August",
		Age:     1,
		Color:   "Beige",
	}
	alex := Person{
		FirstName: "Alex",
		LastName:  "Ferguson",
		Age:       25,
		Sex:       "Male",
		Country:   "UK",
	}

	fmt.Println("August struct: ", august)
	fmt.Println("Alex struct: ", alex)

	var connection Connection
	connection = august
	fmt.Println(connection.Attach(&alex))

}
