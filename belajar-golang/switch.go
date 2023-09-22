package main

import "fmt"

func main() {

	var nama = "Fauzan yanuar putra"
	switch nama {
	case "Fauzan":
		fmt.Println("Hallo Fauzan")
		fmt.Println("Hallo Fauzan")
	case "Anton":
		fmt.Println("Hallo Anton")
		fmt.Println("Hallo Anton")
	case "Kumalasari":
		fmt.Println("Hallo Kumalasari")
		fmt.Println("Hallo Kumalasari")
	default: 
		fmt.Println("kenalan dong")
		fmt.Println("kenalan dong")
	}

	switch len(nama) > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	switch {
	case len(nama) > 10:
		fmt.Println("Nama terlalu panjang")
	case len(nama) > 5:
		fmt.Println("Nama lumayan panjang")
	default: {
		fmt.Println("nama sudah benar")
	}
	}



}