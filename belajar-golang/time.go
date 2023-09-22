package main

import (
	"fmt"
	"time"
)

func main() {
		date := time.Now()

		fmt.Println(date)
		fmt.Println(date.Year())
		fmt.Println(date.Month())
		fmt.Println(date.Day())
		fmt.Println(date.Hour())
		fmt.Println(date.Minute())
		fmt.Println(date.Second())
		fmt.Println(date.Nanosecond())

		utc := time.Date(2023, 10, 10, 10, 10, 10, 10, time.UTC)
		fmt.Println(utc)

		layout := "02-01-2006"
		parse, _ := time.Parse(layout, "03-02-2020")

		fmt.Println(parse)


}