package main

import (
	"fmt"
	"strings"
)

func main() {
		fmt.Println(strings.Contains("Muhammad Fauzan", "Muhammad"))
		fmt.Println(strings.Contains("Muhammad Fauzan", "udin"))
		result := strings.Split("Muhammad Fauzan", " ")
		for _, value := range result {
				fmt.Println( "-",value)
		}
		fmt.Println(strings.ToLower("Muhammad Fauzan Yanuar Putra"))
		fmt.Println(strings.ToUpper("Muhammad Fauzan Yanuar Putra"))

		fmt.Println(strings.Trim("  muhammad fauzan  ", " "))

		fmt.Println(strings.ReplaceAll("Fauzan Muhammad Fauzan", "Fauzan", "Yanuar"))
}