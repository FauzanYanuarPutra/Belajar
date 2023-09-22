package main

import "fmt"

func helloWorld(name string) string {
	result := "Hello "
	if name == "" {
		return result + "Bro"
	} else {
		return result + name
	}
}

func main() {
	// helloWorld("Alex")
	fmt.Println(helloWorld("anjas"))
	fmt.Println(helloWorld(""))

}