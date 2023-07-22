package main

import "fmt"

func tipemap() {

	dataBank := map[string]string{
		"nama":   "haydar",
		"alamat": "citayam",
	}

	dataBank["pekerjaan"] = "mahasiswa"

	//implementasi function map
	fmt.Println(dataBank)
	fmt.Println(dataBank["nama"])
	fmt.Println(dataBank["alamat"])
	fmt.Println(dataBank["pekerjaan"])

	var warung map[string]string = make(map[string]string)
	warung["kerupuk"] = "kerupuk kulit"
	warung["kopi"] = "kapal api"
	warung["sabun"] = "lifebuoy"

	fmt.Println(warung)
	fmt.Println(len(warung))

	//menghapus map

	delete(warung, "sabun")

	fmt.Println(warung)
	fmt.Println(len(warung))

}
