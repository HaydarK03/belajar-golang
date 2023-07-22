package main

import "fmt"

func compare() {

	var nama1 = "haydar"
	var nama2 = "haydar"

	var result bool = nama1 == nama2

	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 < value2)
	fmt.Println(value1 > value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
