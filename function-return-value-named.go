package main

import "fmt"

// value named cukup menuliskan return saja tanpa di ikuti nama variabel nya
func getFullName() (firstName, middleName, lastName string) {
	firstName = "sendhy"
	middleName = "boedhi"
	lastName = "satriya"
	return
}

func main() {
	a, b, c := getFullName()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
