package main

import "fmt"

// cara membuat struct : buat dengan kata kunci type, ikuti dengan nama structnya, kemudian tambahkan struct

type Customer struct {
	Name, Address string
	Age           int
	Married       bool
}

// struct function

func sayHai(customer Customer, name string) {
	fmt.Println("Hai", name, "My name is", customer.Name)
}

func main() {

	// cara menginisialisasi data nya

	var sendhy Customer
	sendhy.Name = "Sendhy"
	sendhy.Address = "solo"
	sendhy.Age = 42
	sendhy.Married = true

	// atau dengan cara

	boedhi := Customer{
		Name:    "sendhy",
		Address: "Solo",
		Age:     42,
		Married: true}

	sayHai(boedhi, "Gaza")

	fmt.Println(sendhy)
	fmt.Println(sendhy.Name)
	fmt.Println(sendhy.Address)
	fmt.Println(sendhy.Age)
	fmt.Println(sendhy.Married)
	fmt.Println(boedhi)
}
