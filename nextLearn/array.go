package main

import "fmt"

func array() {
	var names [4]string
	names[0] = "muhammad"
	names[1] = "haydar"
	names[2] = "ilham"
	names[3] = "kamil"

	//mengakses data array
	fmt.Println(names[2])

	//penggunaan array tanpa deklarasi
	var values = [3]int{
		12,
		13,
		14,
	}
	fmt.Println(values[2])

	//mencari panjang array
	fmt.Println(len(values))
}
