package main

import "fmt"

func main() {
	//func parameter
	sayHello("haydar", "kamil")
	sayHello("budi", "jawa")

	// func return value
	result := getHello("kiwil")
	fmt.Println(result)

	fmt.Println(getHello(""))

	// returning multiple value
	negara, No := returnMV()
	fmt.Println(negara, No)

	// named return value
	kopi, rokok, teh := namedReturnV()
	fmt.Println("kopi =", kopi)
	fmt.Println("rokok =", rokok)
	fmt.Println("teh =", teh)
}
