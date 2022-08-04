package main

import "fmt"

func main() {

	name := "sendhy"

	if name == "sendhy" {
		fmt.Println("Hai sendhy")
	} else if name == "boedhi" {
		fmt.Println("Hello boedhi")
	} else {
		fmt.Println("Hi, nama nya siapa ?")
	}

	// if dengan short statment

	if len(name) > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Ok nama sudah benar")
	}
}
