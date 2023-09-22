package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runAplication(value int){
	defer logging()
	fmt.Println("Run Aplication")
	result := 10 / value
	fmt.Println("result", result)
}


func main() {
	runAplication(0)
	
}