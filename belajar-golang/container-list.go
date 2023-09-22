package main

import (
	"container/list"
	"fmt"
)

func main() {
		data := list.New()
		data.PushBack("Fauzan")
		data.PushBack("Yanuar")
		data.PushBack("Putra")
		data.PushFront("Muhammad")


		// Dari awal sampai akhir
		for e := data.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value)
		}

		// Dari akhir sampai awal
		for e := data.Back(); e != nil; e = e.Prev() {
			fmt.Println(e.Value)
		}

}