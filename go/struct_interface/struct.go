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

func (pet *Pet) Attach(person *Person) string {
	return pet.Name + " is " + person.FirstName + " " + person.LastName + "'s pet."
}

func main() {
	wahyu := Pet{
		Age:     30,
		Species: "Penyu",
	}
	jason := Person{
		FirstName: "Jason",
		LastName:  "Statham",
		Age:       30,
		Sex:       "Non-Binary",
		Country:   "US",
	}
	fmt.Println("Jimmy struct: ", wahyu)
	fmt.Println("Jason struct: ", jason)
	fmt.Println(wahyu.Attach(&jason))
}
