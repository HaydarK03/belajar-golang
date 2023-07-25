package main

import "fmt"

// jika func terlalu panjang pada parameter maka gunakan "type declaration"
type Filter func(string) string

// parameter digunakan sebagai function
func filterWords(word string, filter Filter) {
	wordFiltered := filter(word)
	fmt.Println("your", wordFiltered)
}

func filter(word string) string {
	if word == "bangsat" {
		return "******"
	} else {
		return "good boy " + word
	}
}
