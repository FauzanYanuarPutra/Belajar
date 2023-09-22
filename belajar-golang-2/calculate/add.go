package calculate

import (
	"errors"
	"fmt"
	"reflect"
)


func Add(a int, op rune, b int) (float64, string, error) {
	var result float64
	var operation string
	var err error

	switch op {
	case '+':
		result = float64(a + b)
		operation = "Penjumlahan"
	case '-':
		result = float64(a - b)
		operation = "Pengurangan"
	case 'x':
		result = float64(a * b)
		operation = "Perkalian"
	case ':':
		if b == 0 {
			return 0, "", errors.New("Pembagi tidak boleh nol")
		}
		result = float64(a) / float64(b)
		operation = "Pembagian"
	default:
		err = errors.New("Error: Anda seharusnya memasukkan operator +, -, x, :")
	}

	return result, operation, err
}

func Nilai(score interface{}) {
	var grade interface{}

	if reflect.TypeOf(score).Kind() == reflect.Int {
			sc := score.(int)
			switch{
				case sc <= 50:
					grade = "E"
				case sc <= 60:
					grade = "D"
				case sc <= 70:
					grade = "C"
				case sc <= 100:
					grade = "A"
				default:
					grade = "Null"
			}

			fmt.Println("Nilai Anda: ", grade)
		} else {
			fmt.Println("Masukan Angka!!")
		}

}


func Perulangan(Nama string){
	for i, value := range Nama {
		
		if i % 2 == 0 && value != ' ' {
			// fmt.Println("index ke :", i, " Adalah ", string(value))
			fmt.Println(string(value))
		}

		bosq := string(value)

		switch bosq {
		case "a", "i", "u", "e", "o":
			fmt.Println(string(value), i)
		}
	}
}

func helloWorld() {
	fmt.Println("Hello World! (Private)")
}

func Map(){
	MapData := map[string]string {
		"Javascript": "Belajar Javascript",
		"Golang": "Belajar Golang",
		"PHP": "Belajar PHP",
		"C++": "Belajar C++",
	}

	for key, value := range MapData {
		fmt.Println("key : ", key, " Value : ", value)
	}

	fmt.Println("=========")
	delete(MapData, "Golang")

	for key, value := range MapData {
		fmt.Println("key : ", key, " Value : ", value)
	}


}