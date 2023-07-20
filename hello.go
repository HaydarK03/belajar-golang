package main

import (
	"fmt"
)

func main() {
	//ini hello
	fmt.Println("hello worlds")

	//dibawah ini tipe data

	//ini number
	fmt.Println("satu =", 1)
	fmt.Println("dua =", 2)

	//ini boolean
	fmt.Println("benar =", true)
	fmt.Println("salah =", false)

	//ini string
	fmt.Println(len("haydar"))
	fmt.Println(len("haydar ilham"))
	fmt.Println("haydar ilham kamil"[4])

	//ini variable
	var nama string

	nama = ("haydar")
	fmt.Println(nama)

	nama = ("kamil")
	fmt.Println(nama)

	//optimize var tanpa perlu menyebutkan tipe data
	var age = 30
	fmt.Println(age)

	//optimize var menggunakan :=
	negara := ("indonesia")
	fmt.Println(negara)

	//deklarasi multiple variable
	var (
		FirstName = ("Haydar")
		lastName  = ("Kamil")
	)
	fmt.Println(FirstName)
	fmt.Println(lastName)

	//ini constant
	const (
		namadepan    = "haydar"
		namabelakang = "kamil"
		value        = 1080
	)
	fmt.Println(namadepan)
	fmt.Println(namabelakang)
	fmt.Println(value)

	//konversi tipe data
	var nilai32 int32 = 100080
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	//konversi tipe data string
	var saya = "haydar"
	var e byte = saya[0]
	var estring string = string(e)

	fmt.Println(saya)
	fmt.Println(estring)
}
