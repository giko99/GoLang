package main

import "fmt"

func main() {
	fmt.Println("Hello World", "now I'am learning", "GO Languange")

	var firstName string = "Gigih"

	var lastName string
	lastName = "Pambuko"

	//hasilnya akan sama pada ketiga cara menampilkan dibawah
	fmt.Printf("Hello %s %s!\n", firstName, lastName)

	//Println otomatis menambahkan baris baru
	fmt.Println("Hello", firstName, lastName+"!")

	fmt.Printf("halo Gigih Pambuko!\n")

	var namaDepan string = "John"

	//penggunaan ":=" hanya sekali diawal, seterrusnya bisa menggunakan "=" langsung
	namaBelakang := "Wick"

	namaBelakang = "Ethan"

	fmt.Printf("Hello %s %s!\n", namaDepan, namaBelakang)

	//deklarasi multi variable
	one, isMonday, twoPointTwo, say := 1, true, 2.2, "Hello"

	fmt.Println(one, isMonday, twoPointTwo, say)

	//menyimpan variable yg tidak digunakanmenggunakan karakter "_"
	_ = "belajar Golang"
	_ = "Golang itu mudah"
	nama, _ := "Ari", "Ester"
	fmt.Println(nama)

	name := new(string)
	fmt.Println(name)  //0xc0000102c0
	fmt.Println(*name) // ""

	//tipe data numerik Non-Desimal
	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644

	fmt.Printf("bilangan positif : %d\n", positiveNumber)
	fmt.Printf("bilangan negatif : %d\n", negativeNumber)

	//tipe data numerik desimal
	var decimalNumber = 2.62

	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)

	//tipe data boolean
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)

	//tipe data string
	var message = `Nama saya "Giko".
Salam kenal.
Mari belajar "Golang".`

	fmt.Println(message)

	//penggunaan konstanta seperti pi(22/7), kecepatan cahaya (299.792.458 m/s)
	const sapaGue string = "Giko"
	fmt.Print("Halo ", sapaGue, "!\n")

	fmt.Println("john wick")
	fmt.Println("john", "wick")

	fmt.Print("john wick\n")
	fmt.Print("john ", "wick\n")
	fmt.Print("john", " ", "wick\n")
}
