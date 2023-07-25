package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil func")
}

// defer akan menghiraukan error dan akan lanjut ke perintah selanjutnya
func runApp(value int) {
	defer logging()
	fmt.Println("run application")
	result := 10 / value
	fmt.Println("result ", result)

}
