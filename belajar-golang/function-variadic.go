package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}

func main() {
	a := sumAll(2, 342, 42, 523, 5235, 13)
	fmt.Println(a)

	numbers := []int {13,123,234234,2342342,32423423}

	bTotal := sumAll(numbers...)
	fmt.Println(bTotal) 
}