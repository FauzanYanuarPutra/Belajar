package main

import (
	"belajar-golang-2/calculate"
	"fmt"
)

type User struct {
	ID int
	FirstName string
	LastName string
	Email string
	IsActive bool
}

type Student struct {
	ID int
	Name string
	GPA float32
}

func (student *Student) Graduate(){
	student.Name = student.Name + " S.T"
}

func (user User) displayUser() interface{} {
	return "Halllo " + user.FirstName + ", Apakah Benar Ini Email Anda ? -" + user.Email
}

type Group struct {
	name string
	Admin User
	Users []User
	IsAvailable bool
}


func (group Group) displayGroup() {
	fmt.Printf("Group Name: %s\n", group.name)
	fmt.Printf("Admin: %v\n", group.Admin.FirstName)
	fmt.Printf("Members%v\n", ":")
	for _, value := range group.Users {
		fmt.Printf("%v\n", value.FirstName)
	}
	fmt.Printf("IsActive: %v\n", group.IsAvailable)
}

type Gamer struct {
	Name string
	Games []string
}

func (gamer *Gamer) AddGame(Name string) {
		gamer.Games = append(gamer.Games, Name)
		// fmt.Println(Name)
		// gamer.Games = append(gamer.Games, "Mario")
}


type BangunDatar interface {
		HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int{
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

func (persegipanjang PersegiPanjang) HitungLuas() int{
	return persegipanjang.Panjang * persegipanjang.Lebar
}

type Asal struct {
	Angka int
}

func (asal Asal) HitungLuas() int {
	return 22222
}

func main() {
		// result := entity.TestAja("Fauzan Yanuar Putra", 20)
		// result2, operator, err  := calculate.Add(20, 's', 20)
		// if err != nil {
		// 	fmt.Println(err.Error())
		// } else {
		// 	fmt.Println(result2, operator)
		// }

		// var name string

		// name = "Fauzan"
		// age := 10

		// fmt.Println(name)

		// fmt.Println(result2)

		// fmt.Println(result)
		// fmt.Println(calculate.ShowHelloWorld("hallo"))

		// fmt.Println(age)
		// if age >= 10 {
		// 	fmt.Println("Boleh Bermain Game")
		// } else {
		// 	fmt.Println("Tidak Boleh Bermain Game")
		// }

		// calculate.Nilai(50)

		// calculate.Perulangan("Hallo Saya Muhammad Fauzan")

		// calculate.Map()


		calculate.Kuis1()

		// scores := []int{10,5,8,9,7, 20}

		// total := sum(scores)

		// fmt.Println(total)

		fmt.Println("Struct")

		var ShowUser []User

		AddUser(&ShowUser, "Fauzan", "Yanuar", "Fauzan@gmail.com", true)
		AddUser(&ShowUser, "Sasuke", "Sasuke", "Susanoo@gmail.com", true)
		AddUser(&ShowUser, "Anjas", "Sasuke", "Susanoo@gmail.com", true)
		AddUser(&ShowUser, "asdas", "Sasuke", "Susanoo@gmail.com", true)
		AddUser(&ShowUser, "wwewe", "Sasuke", "Susanoo@gmail.com", true)
		AddUser(&ShowUser, "12312", "Sasuke", "Susanoo@gmail.com", true)

		var ShowGroup []Group


		users := []User{ShowUser[0], ShowUser[1]}
		users2 := []User{ShowUser[4], ShowUser[3]}


		ShowGroup = append(ShowGroup, Group{"prolov", ShowUser[0], users, true})
		ShowGroup = append(ShowGroup, Group{"Anjay", ShowUser[3], users2, false})

		

		for _, user := range ShowUser {
			fmt.Println(user.displayUser())
		}

		for i, group := range ShowGroup {
			fmt.Println(i + 1)
			group.displayGroup()
		}

		number := 5

		fmt.Println(number)

		ChangeNumber(&number, 2000)

		fmt.Println(number)


		student := Student{1, "Muhammad Fauzan Yanuar Putra", 3.40}

		fmt.Println(student.Name)


		student.Graduate()


		fmt.Println(student.Name)


		gamer := Gamer{Name: "zelda"}

		

		gamer.AddGame("Mario")
		gamer.AddGame("fifa 2020")
		gamer.AddGame("Soccer 2020")

		fmt.Println("Nama Game :", gamer.Name)
		fmt.Println("Game Apa Saja")

		for _, game := range gamer.Games {
			fmt.Println(game)
		}

		persegi := Persegi{Sisi: 35}
		fmt.Println("Persegi : ",persegi.HitungLuas())

		persegiPanjang := PersegiPanjang{Panjang: 20, Lebar: 30}
		fmt.Println("Persegi Panjang : " ,persegiPanjang.HitungLuas())

		asal := Asal{Angka: 20}
		fmt.Println("Asal : " ,asal.HitungLuas())


}




func ChangeNumber(old *int, new int){
	*old = new
}

func AddUser(users *[]User, firstName, lastName,  email string, isActive bool) {
	user := User {
		ID: len(*users) + 1,
		FirstName: firstName,
		LastName: lastName,
		Email: email,
		IsActive: isActive,
	}

	*users = append(*users, user)
}





func sum(scores []int) int {
	score := 0
	for _, value := range scores { 
		score += value
	}

	return score
}
