package main

import (
	"fmt"
)

func random() interface{}{
	return true
}

func main() {
	result := random()
	// data := result.(string)
	// fmt.Println(data)

	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is String")
	case int:
		fmt.Println("Value", value, "is Int")
	default: 
		fmt.Println("Unknown Type")
	}
}