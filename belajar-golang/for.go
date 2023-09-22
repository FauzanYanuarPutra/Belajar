package main

import "fmt"

func main() {

	// counter := 1
	// for counter <= 10 {
	// 	fmt.Println(counter)
	// 	counter++
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// slice := []string{"Muhammad", "Fauzan", "Yanuar", "Putra", "Masasih"}

	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println(slice[i])
	// }

	// for i, value := range slice {
	// 	fmt.Println(i, value)
	// }

	// for _, value := range slice {
	// 	fmt.Println(value)
	// }

	person := make(map[string]string)
	person["nama"] = "Fauzan"
	person["address"] = "Bandung"

	for key, value := range person {
		fmt.Println(key, value)
	}

}