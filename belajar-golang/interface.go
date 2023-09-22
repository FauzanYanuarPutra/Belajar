package main

import "fmt"

type HasName interface {
	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hallo", hasName.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName()string {
	return person.Name
}

type Animal struct {
	Name string
}

func (a Animal) GetName()string {
	return a.Name
}

func main() {
	fauzan := Person{}
	fauzan.Name = "Fauzan"

	sayHello(fauzan)

	cat := Person{
		Name: "cat",
	}

	sayHello(cat)
	// fmt.Println(fauzan)
}