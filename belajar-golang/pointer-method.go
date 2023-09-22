package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married(){
	man.Name = "Mr. " + man.Name
}



func main() {
		fauzan := Man{"Fauzan"}
		fauzan.Married()
		fmt.Println(fauzan.Name)
}