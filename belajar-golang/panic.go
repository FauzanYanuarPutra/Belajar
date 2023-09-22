package main

import "fmt"

func endApp(){
	fmt.Println("End app")
	message := recover()
	if message != nil {
		fmt.Println("message", message)
	}
}


func runApp(error bool){
	defer endApp()
	if error {
		panic("Error")
	}
	fmt.Println("aplikasi berjalan lancar")
}

func main() {
		runApp(true)
		fmt.Println("Fauzan")
}