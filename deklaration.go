package main

import "fmt"

func declaration() {
	type NoKTP string
	type Status bool

	var ktphaydar NoKTP = "112122314"
	var kawin Status = false
	fmt.Println(ktphaydar)
	fmt.Println(kawin)
}
