package main

import "fmt"

func main() {
	//dasar
	fmt.Println("Hello World!", "how", "are", "you", "today", "?")
	//deklarasi variabel menggunakan keyword var
	var firstName string = "saiful"

	var lastName string
	lastName = "arif"

	fmt.Printf("halo %s %s! \n", firstName, lastName)
	fmt.Println("halo " + firstName + " " + lastName + "!")
	//cara lain
	coba1 := "tes1" //tidak perlu menambahkan var atau string
	coba2 := "tes2"
	var coba3 string = "tes3"

	fmt.Println(coba1 + " " + coba2 + " " + coba3)
	fmt.Printf("bro %s %s %s \n", coba1, coba2, coba3)
	fmt.Println("bro %s %s %s \n", coba1, coba2, coba3) //ada perbedaan dalam penggunaan printf dan println

	//deklarasi multi variabel
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"

	fmt.Printf("halo %s %s %s %s %s %s %s %s %s %s\n", first, second, third, seventh, eight, ninth, one, isFriday, twoPointTwo, say)

	//variabel underscore _
	_ = "belajar golang"

	//deklarasi variabel menggunakan keyword new
	name := new(string)
	fmt.Println(name)
	fmt.Println(*name)
}

//ini komentar cuy
/*
ini juga
*/
