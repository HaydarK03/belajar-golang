package main

import "fmt"

func closure() {
	name := "haydar"
	number := 1
	// space yang berada dibawah dapat mempengaruhi space diatasnya
	increment := func() {
		name = "joko"

		number++

	}
	increment()

	fmt.Println(name)
	fmt.Println(number)
}
