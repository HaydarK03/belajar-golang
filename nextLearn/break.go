package main

import "fmt"

// break akan menghentikan perulangan ke jumlah yang ditentukan
func breaks() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("hitungan ke", i)
	}
}
