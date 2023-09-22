package helper

import "fmt"

var Application = "Golang"
var versi = 23


func SayHello(nama string) {
	fmt.Println("Hallo", nama)
}

func sayGoodbye(nama string){
	fmt.Println("Bye", nama)
}