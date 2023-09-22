package main

import "fmt"

type Customer struct {
	Name, Adress string
	Age int
}

func (customer Customer) sayHello(){
	fmt.Println("Hallo, My name is", customer.Name)
}

func main() {
		fauzan := Customer{}
		fauzan.Name = "fauzan"
		fauzan.Adress = "Bandung"
		fauzan.Age = 19

		fmt.Println(fauzan)
		fmt.Println(fauzan.Name)
		fmt.Println(fauzan.Adress)
		fmt.Println(fauzan.Age)

		jomba := Customer{
			Name: "Jomba",
			Adress: "Jakarta",
			Age: 12,
		}

		fmt.Println(jomba)
		jomba.sayHello()


		// not recomended
		norec := Customer{"norec", "tangerang", 20}
		fmt.Println(norec)
		norec.sayHello()




}