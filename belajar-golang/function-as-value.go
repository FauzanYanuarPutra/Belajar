package main

import "fmt"

func getGoodbye(Nama string) string {
	return "Bye " + Nama
}

func main() {
	sayGoodBye := getGoodbye

	result := sayGoodBye("Fauzan")
	fmt.Println(result)
	fmt.Println(sayGoodBye("var"))
	fmt.Println(getGoodbye("func"))
}