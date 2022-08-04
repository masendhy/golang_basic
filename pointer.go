package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	// di golang passing data itu by value, jadi saat kita rubah data di address1 saja address2 tidak ikut berubah, demikian pula sebaliknya. Untuk merubah menjadi passing data by reference gunakan pointer, yaitu dengan menggunakan  operator & yaitu dengan menambahkan tanda baca & di depan data asignment nya.

	//pass by value
	address1 := Address{"Solo", "Central Java", "Indonesia"}
	address2 := address1

	address1.City = "Cepu"

	//pass by reference
	address3 := &address2

	address3.Province = "Jawa Tengah"

	// menggunakan operator * untuk merubah keseluruhan data dari data  awal

	*address3 = Address{"Malang", "East Java", "Indonesia"}

	//menggunakan kata kunci new untuk mengkosongkan data

	address3 = new(Address)

	// mengisi data variable

	address3.City = "Jogja"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

}
