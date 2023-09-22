package main

import "fmt"

type Address struct {
	Kota, Provinsi, Negara string
}

func addNegara(address *Address) {
	address.Negara = "Bandung"
}

func main() {
		address1 := Address{"Bandung", "Jawa Barat", "Indonesia"}
		address4 := &Address{"Jakarta", "Jawa Barat", "Indonesia"}
		address2 := &address1
		address3 := &address1

		*address2 = Address{"INi", "DAta", "SAlah"}

		fmt.Println(address1)
		fmt.Println(address2)
		fmt.Println(address3)
		fmt.Println(address4)

		address5 := new(Address)
		address5.Kota = "Assalamualaikum"

		fmt.Println(address5)

		alamat := Address{
			Kota: "Maret",
			Provinsi: "Maret",
		}

		addNegara(&alamat)
		fmt.Println(alamat)


}