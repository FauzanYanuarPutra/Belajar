package main

import "fmt"

func main() {
		nama := "Budi"
		counter := 0

		increment := func() {
			counter := 0
			nama := "Fauzan"
			counter++
			fmt.Println(nama)
			fmt.Println(counter)
		}

		increment()


		fmt.Println(nama)
		fmt.Println(counter)
		// fmt.Println(increment)

}