package main

import (
	"fmt"
	"regexp"
)

func main() {
		regex := regexp.MustCompile("e([a-z])o")

		fmt.Println(regex.MatchString("eko"))
		fmt.Println(regex.MatchString("eto"))
		fmt.Println(regex.MatchString("eDo"))
		fmt.Println(regex.MatchString("eba"))

		search := regex.FindAllString("eko eta emo omo",  -1)
		fmt.Println(search)

}