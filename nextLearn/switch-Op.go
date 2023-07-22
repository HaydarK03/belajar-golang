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

	//switch shorts statement
	// switch length := len(customer); length > 13 {
	// case true:
	// 	fmt.Println("apakah anda ingin membuat member?")
	// case false:
	// 	fmt.Println("terima kasih sudah menjadi member")
	// }

	length := len(customer)
	switch {
	case length > 13:
		fmt.Println("apakah anda ingin membuat member?")
	case length == 12:
		fmt.Println("terima kasih sudah menjadi member")
	default:
		fmt.Println("stand by")
	}
}
