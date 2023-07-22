package main

import "fmt"

func ifOP() {
	var name = "haydar"

	if name == "haydar" {
		fmt.Println("hay haydar")

	} else {
		fmt.Println("kamu siapa")
	}

	//if else expresion
	nama := "kamil"

	if nama == "kamil" {
		fmt.Println("hay kamil")
	} else if nama == "haydar" {
		fmt.Println("hay haydar")
	} else {
		fmt.Println("lu siapa dah?")
	}

	//if short statemen
	if length := len(nama); length >= 6 {
		fmt.Println("pas mantap")
	} else {
		fmt.Println("nama lu kepanjangan")
	}
}
