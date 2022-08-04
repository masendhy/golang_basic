package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// dari pada harus membuat function blacklistnya berulangkali untuk setiap user

// func blacklistAdmin(name string)bool {
// 	return name == "admin"
// }

// func blacklistRoot(name string) bool {
// 	return name == "root"
// }

func main() {

	// maka lebih baik dibuatkan variabel dengan anonymous function

	blacklist := func(name string) bool {
		return name == "admin" // name untuk yang di blacklist
	}

	registerUser("admin", blacklist)
	registerUser("sendhy", blacklist)

}
