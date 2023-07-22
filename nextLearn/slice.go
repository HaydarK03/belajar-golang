package main

import "fmt"

func slice() {
	var bulan = [...]string{
		"muharam",
		"safar",
		"rabiul awal",
		"rabiul akhir",
		"jumadil awal",
		"jumadil akhir",
		"rajab",
		"syakban",
		"ramadhan",
		"syawal",
		"zulkaidah",
		"zulhijah",
	}
	var slice1 = bulan[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//append
	//hati-hati saat menggunakan append "data array akan terganti ketika kapasitas memadai"
	var slice2 = bulan[10:]
	fmt.Println(slice2)

	slice3 := append(slice2, "maret")
	fmt.Println(slice3)

	fmt.Println(bulan)

	//make slice
	newSlice := make([]string, 2, 10)
	newSlice[0] = "haydar"
	newSlice[1] = "kamil"

	fmt.Println(newSlice)
	fmt.Println(cap(newSlice))

	//copy slice *kapasitas array harus sama
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//hati hati dengan menggunakan slice
	iniarray := [...]int{1, 2, 3, 4, 5}
	inislice := []int{1, 2, 3, 4, 5}
	fmt.Println(iniarray)
	fmt.Println(inislice)
}
