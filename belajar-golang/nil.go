package main

import (
	"fmt"
	"strings"
)

func NewMap(nama string) map[string]string {
	if strings.TrimSpace(nama) == "" {
		return nil
	} else {
		return map[string]string{
			"nama": strings.TrimSpace(nama),
		}
	}
}

func main() {
	person := NewMap("  ")
	if person == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(person["nama"])
	}
}
