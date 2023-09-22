package main

import "fmt"

func main() {
	person := map[string]string{
		"nama":    "Fauzan",
		"address": "Bandung",
	}

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["address"])
	fmt.Println(len(person))

	var book map[string]string = make(map[string]string)

	book["nama"] = "Alex"
	book["address"] = "Jakarta"
	book["salah"] = "hapuskan"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "salah")

	fmt.Println(book)
	fmt.Println(len(book))







}