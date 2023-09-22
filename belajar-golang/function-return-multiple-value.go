package main

import "fmt"

func getFullName() (string, string, string) {
	return "muhammad", "Fauzan", "Yanuar"
}

func main() {
	firstName, _, _ := getFullName()
	fmt.Println(firstName)
	// fmt.Println(midleName)
	// fmt.Println(lastName)

}