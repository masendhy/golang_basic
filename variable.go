package main

import "fmt"

func main() {

	// membuat variabel dengan menyebutkan type datanya
	var name string

	name = "masendhy"
	fmt.Println(name)

	name = "abik"
	fmt.Println(name)

	// variabel tanpa type data
	var color = "white"
	fmt.Println(color)

	// mengganti kata kunci var dengan tanda := untuk deklarasi awal
	city := "solo"
	fmt.Println(city)

	// multiple variable
	var (
		firstName = "sendhy"
		lastName  = "satriya"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
