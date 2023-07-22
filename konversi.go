package main

import (
	"fmt"
)

func tipekonversi() {
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
