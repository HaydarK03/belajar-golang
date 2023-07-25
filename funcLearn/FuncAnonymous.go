package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("you are blocked", name)
	} else {
		fmt.Println("welcome", name)
	}

}

// func blacklistAdmin(name string)bool{
// return name == "root"
// }
//
// func blacklistRoot(name string)bool{
// return name == "root"
// }
