package main

import "fmt"

// menambahkan parameter didalam function harus diikuti dengan type data nya

func sayHai(firstName string, lastName string) {
	fmt.Println("Hai", firstName, lastName)
}
func main() {
	sayHai("sendhy", "boedhi")

	firstName := "masendhy"
	lastName := "bagus"
	sayHai(firstName, lastName)

}
