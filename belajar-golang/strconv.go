package main

import (
	"fmt"
	"strconv"
)

func main() {

		// string ke bool
		stringToBool, err:= strconv.ParseBool("true")
		if err == nil {
			fmt.Println(stringToBool)
		} else {
			fmt.Println(err.Error())
		}

		// string ke int
		stringToInt, err := strconv.ParseInt("200000", 10, 64)
		if err == nil {
			fmt.Println(stringToInt)
		} else {
			fmt.Println(err.Error())
		}

		// atoi itoa
		// atoi itu dari string ke int
		atoi, _ := strconv.Atoi("13213213123") 
		fmt.Println(atoi)

		// itoa itu dari int ke string
		itoa := strconv.FormatInt(20000000, 10) 
		fmt.Println(itoa)
}