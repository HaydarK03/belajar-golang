package main

import "fmt"

func ifOP() {
	var name = "haydar"

	if name == "haydar" {
		fmt.Println("hay haydar")

	} else {
		fmt.Println("kamu siapa")
	}

	//if else statemen
	nama := "kamil"

	if nama == "kamil" {
		fmt.Println("hay kamil")
	} else if nama == "haydar" {
		fmt.Println("hay haydar")
	} else {
		fmt.Println("lu siapa dah?")
	}

	if length := len(nama); length >= 6 {
		fmt.Println("pas mantap")
	} else {
		fmt.Println("nama lu kepanjangan")
	}
}
