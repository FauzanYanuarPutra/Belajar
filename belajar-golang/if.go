package main

import "fmt"

func main() {
	var nama = "Fauzan"
	var umur = 10


	if nama == "Fauzan" && umur == 10 {
		fmt.Println("Hallo Fauzan")
	} else if nama == "Anton" {
		fmt.Println("Hallo Anton")
	} else if nama == "Jota" {
		fmt.Println("Hallo Jota")
	} else {
		fmt.Println("Hai, Kenalan dong!")
	}


	if len(nama) > 5 {
		fmt.Println("Nama Terlalu Panjang")
	}

}