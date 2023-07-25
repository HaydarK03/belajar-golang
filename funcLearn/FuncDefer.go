package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil func")
}

func runApp(value int) {
	fmt.Println("run application")
	result := 10 / value
	fmt.Println("result ", result)
	logging()
}
