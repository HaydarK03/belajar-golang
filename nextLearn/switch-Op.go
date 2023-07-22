package main

import "fmt"

func switchOp() {
	customer := "old customer"

	switch customer {
	case "new customer01":
		fmt.Println("apakah ada yang bisa dibantu?")
	case "old customer":
		fmt.Println("terima kasih telah menjadi pelanggan setia")
	default:
		fmt.Println("stand by")
	}

	// fmt.Println(len(customer))
	switch length := len(customer); length > 13 {
	case true:
		fmt.Println("apakah anda ingin membuat member?")
	case false:
		fmt.Println("terima kasih sudah menjadi member")
	}
}
