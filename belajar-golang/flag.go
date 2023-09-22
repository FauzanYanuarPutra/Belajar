package main

import (
	"flag"
	"fmt"
)

func main() {
		host := flag.String("host", "localhost", "masukan localhost nya")
		username := flag.String("username", "root", "masukan username nya")
		password := flag.String("password", "root", "masukan password nya")
		number := flag.Int("number", 123, "masukan password nya")
		

		flag.Parse()

		fmt.Println(*host)
		fmt.Println(*username)
		fmt.Println(*password)
		fmt.Println(*number)



}