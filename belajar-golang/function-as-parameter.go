package main

import (
	"fmt"
	"strings"
)

type Filter func(string)string

func sayHelloWithFilter(nama string, filter Filter){
	fmt.Println("Hello ", filter(nama))
}

func spamFilter(nama string)string {
	if nama == "Anjing" {
		return "..."
	} else {
		return nama
	}
}

func upperFilter(nama string)string {
	return strings.ToUpper(nama)
}

func main() {
	sayHelloWithFilter("fauzan", upperFilter)
	sayHelloWithFilter("Anjing", spamFilter)

}