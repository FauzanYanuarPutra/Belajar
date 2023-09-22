package main

import "fmt"

type BlackList func(string) bool

func registerUser(nama string, blackList BlackList) {
	if blackList(nama) {
		fmt.Println("You are blocked", nama)
	} else {
		fmt.Println("Welcome", nama)
	}
}


func main() {
	blacklist := func(nama string) bool {
		return nama == "admin" || nama == "fauzan"
	}

	registerUser("admin", blacklist)
	registerUser("kasir", blacklist)
	registerUser("fauzan", blacklist)

	registerUser("golang", func(nama string)bool {
		return nama == "golang"
	})

}