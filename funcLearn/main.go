package main

import (
	"fmt"
)

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

	// variadic function
	total := sumAll(10, 20, 30, 10)
	fmt.Println(total)

	slice := []int{1, 2, 3, 4, 5}
	total = sumAll(slice...)
	fmt.Println(total)

	//function value
	goodbye := getGoodBye
	fmt.Println(goodbye("haydar"))

	//function as parameters
	filterWords("test", filter)
	filterWords("bangsat", filter)

	// function anonymous
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("haydar", blacklist)

	// Anon func ditambahkan langsung dengan parameter
	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("haydar", func(name string) bool {
		return name == "root"
	})

	//fucntion recursive
	loop := factorialLoop(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)
}
