package main

import "fmt"

func main(){
	var nama [3]string

	nama[0] = "Muhammad"
	nama[1] = "Fauzan"
	nama[2] = "Yanuar"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	var values = [3]int{
		98,
		80,
		20,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	fmt.Println(len(values))
	fmt.Println(len(nama))


	var lagi [10]int

	fmt.Println(len(lagi))

}