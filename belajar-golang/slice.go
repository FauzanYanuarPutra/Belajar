package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//months[5] = "Diubah"
	//fmt.Println(slice1)

	//slice1[0] = "Mei Lagi"
	//fmt.Println(months)

	var slice2 = months[11:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Eko")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 9)

	newSlice[0] = "new"
	newSlice[1] = "slice"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlite := make([]string, len(newSlice), cap(newSlice))
	copy(copySlite, newSlice)
	fmt.Println(copySlite)

	var iniArray = [...]int{
		1,2,3,4,5,
	}

	var iniSlice = []int{
		1,2,3,4,5,
	}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)


}