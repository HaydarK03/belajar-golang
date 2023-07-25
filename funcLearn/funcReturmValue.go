package main

func getHello(name string) string {

	// return harus mengembalikan data yang sama dengan return value
	// kode program setelah return tidak akan dieksekusi
	if name == "" {
		return "hai guys"

	} else {
		return "hello " + name
	}

}
