package main

import "fmt"

//bikin kontrak baru dalam bentuk interface

type HasName interface {
	// kita ingin jika mau ngikutin kontak ini dia harus mempunyai function GetName dengan return string

	GetName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

//pertama pertama kita bikin struck
type Person struct {
	Name string
}

// kemudian bikin function nya
func (person Person) GetName() string {
	return person.Name
}

// bikin struck yang beda
type Animal struct {
	Name string
}

// ini function dr struck animal
func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var sendhy Person
	sendhy.Name = "sendhy"

	sayHello(sendhy)

	cat := Animal{
		Name: "Pussy",
	}

	sayHello(cat)
}
