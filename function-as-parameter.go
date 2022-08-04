package main

import "fmt"

// type declaration untuk mempersingkat penulisan string di dalam parameter

// ex :
// type Filter func(string)string
// jadi pada function filternya tinggal menuliskan
// func sayHelloWithFIlter(name(string),filter Filter)string {}

func sayHelloWithFilter(name string, filter func(string) string) {

	filteredName := filter(name)
	fmt.Println("hello", filteredName)
}

// kita buat function filternya yang akan kita jadikan filter function
func spamFilter(name string) string {
	if name == "Anjing" {
		return "zzzzz"
	} else {
		return name
	}
}
func main() {
	sayHelloWithFilter("masendhy", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
