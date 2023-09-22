package main

import "fmt"

func helloWorld(firstName string, lastName string) {
	fmt.Println("Hallo", firstName, lastName)
}

func main(){
	firstName := "Muhammad"
	helloWorld(firstName, "Fauzan")
	helloWorld("Yanuar", "Putra")

}