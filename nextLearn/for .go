package main

import "fmt"

func forLoops() {
	hitung := 1

	for hitung <= 10 {
		fmt.Println("hitungan ke", hitung)
		hitung++
	}

	// for dengan statement

	// init statement % post statement
	for hitung := 1; hitung <= 10; hitung++ {
		fmt.Println("hitungan ke", hitung)
	}

	// akses data slice atau array secara manual

	slice := []string{"haydar", "ilham", "kamil", "budi", "joko"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	for i, value := range slice {
		fmt.Println("index", i, "=", value)
	}

	orang := make(map[string]string)
	orang["ID"] = "00001"
	orang["Name"] = "Alex"

	for key, value := range orang {
		fmt.Println(key, "=", value)
	}
}
