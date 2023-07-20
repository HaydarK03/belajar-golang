package main

import (
	"fmt"
)

func variable() {
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
}
