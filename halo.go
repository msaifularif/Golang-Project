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
	//Tipe Data
	var positiveNumber uint8 = 90
	var negativeNumber = -1243423644
	fmt.Printf("bilangan positif: %d \n", positiveNumber)
	fmt.Printf("bilangan negatif: %d \n", negativeNumber)
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f \n", decimalNumber)
	fmt.Printf("bilangan desimal: %.10f \n", decimalNumber)
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)
	var baru = `cihuy
	suka-suka lah "Coba"
	iyeeee`
	fmt.Printf("exist? %s \n", baru)
	const nama1 = "gue"
	fmt.Print("lapar ", nama1)
	fmt.Println("john wick")
	fmt.Println("john", "wick")
	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")
	var value = ((2+4)*5/2 - 2)
	fmt.Print(value)
	var ope = ((8%3)*2 + 4)
	var op1 = (ope == 7)
	fmt.Print("Nilai: %d %t \n", ope, op1)
	var left = false
	var right = true
	var leftAndRight = left && right
	fmt.Printf("Nilai: \t(%t) \n", leftAndRight)
	var rightOrLeft = left || right
	fmt.Printf("Nilai: \t(%t) \n", rightOrLeft)
	var notRight = !right
	fmt.Printf("Nilai: \t(%t) \n", notRight)
	var point = 7
	if point == 10 {
		fmt.Println("Kamu lulus sempurna")
	} else if point > 6 {
		fmt.Println("Kamu lulus oke")
	} else {
		fmt.Println("kamu harus remedial deh kayaknya")
	}
	var pointLain = 6840.0
	percent := pointLain / 100
	if percent >= 100 {
		fmt.Printf("%1.f%s perfect! \n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%1.f%s good! \n", percent, "%")
	} else {
		fmt.Printf("%1.f%s not bad! \n", percent, "%")
	}
	var cobaCase = 1
	if cobaCase >= 7 {
		switch cobaCase {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}

	} else {
		if cobaCase == 6 {
			fmt.Println("lumayan lah")
		} else if cobaCase == 5 {
			fmt.Println("jelek loh")
		} else {
			fmt.Println("belajar woi!")
		}

	}

}

//ini komentar cuy
/*
ini juga
*/
