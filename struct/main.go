package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// First, we declare a person with omitempty: if the value has the zero
// value, it is ommitted.
type Person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {

	// We declare a person with all attributes
	person1 := Person{
		Name: "Jean",
		Age:  10,
	}

	// We declare another a person without all attributes. Attributes
	// will take the zero value.
	person2 := Person{}

	// We print person1
	fmt.Printf("Person1: %+v\n", person1)
	fmt.Println("Another way to print person1: ", person1)
	fmt.Printf("Person2: %+v\n", person2)
	jsonValuePerson1, err := json.Marshal(person1)

	if err != nil {
		log.Fatal("cannot marshall value: ", person1)
	}

	fmt.Println("JSON value Person 1: ", string(jsonValuePerson1))

	// We marshall person2 into JSON
	jsonValuePerson2, err := json.Marshal(person2)

	if err != nil {
		log.Fatal("cannot marshall value: ", 2)
	}

	// no element are showing
	fmt.Println("JSON value Person 2: ", string(jsonValuePerson2))

	var decodedValue Person
	err = json.Unmarshal(jsonValuePerson1, &decodedValue)
	if err != nil {
		log.Fatal("error while decoding")
	}
	fmt.Printf("Decoded Value Person 1: %+v\n", decodedValue)
}
