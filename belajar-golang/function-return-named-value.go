package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Muhammad"
	middleName = "Fauzan"
	lastName = "Yanuar"

	return

}

func main() {
	a, _, c := getFullName()
	fmt.Println(a)
	fmt.Println(c)

	// fmt.Println(midleName)
	// fmt.Println(lastName)

}