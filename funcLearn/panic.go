package main

import "fmt"

// jika terdapat sistem error, panic digunakan untuk menghentikan program
// recover digunakan untuk mengambil data dari panic, recover akan dijalankan ketika program error
func endApp() {
	massage := recover()
	if massage != nil {
		fmt.Println("ERROR BRO")
	}
	fmt.Println("aplikasi selesai")
}

func runApps(error bool) {
	defer endApp()
	if error {
		panic("404 NOT FOUND")
	}
	fmt.Println("aplikasi berjalan")
}
