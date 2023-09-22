package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return  0, errors.New("Pembagi Tidak Boleh bernilai 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
}

func main() {
		nilai, err := pembagi(100,0)
		if err == nil {
			fmt.Println(nilai)
		} else {
			fmt.Println("Error", err.Error())
		}

}