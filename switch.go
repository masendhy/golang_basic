package main

import "fmt"

func main() {
	name := "sendhy"

	switch name {
	case "sendhy":
		fmt.Println("Hai", name)
	case "boedhi":
		fmt.Println("Hello", name)
	default:
		fmt.Println("Hai, kenalan dong")
	}

	// switch short statement
	switch len(name) > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}
}
