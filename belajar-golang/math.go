package main

import (
	"fmt"
	"math"
)

func main() {
		fmt.Println(math.Round(2.3))
		fmt.Println(math.Floor(2.8))
		fmt.Println(math.Ceil(2.2))

		fmt.Println(math.Min(2.2, 30.3))
		fmt.Println(math.Max(2.2, 31.3))
}