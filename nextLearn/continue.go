package main

import "fmt"

// continue akan mengskip perulangan selanjutnya namun post statement tetap dijalankan
func continus() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
