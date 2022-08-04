package main

import (
	"fmt"
)

func main() {

	person := map[string]string{ // membuat map baru
		"name":    "masendhy",
		"address": "solo",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	person["age"] = "42"
	fmt.Println(person)
	fmt.Println(person["age"])

	// Function Map
	// len(map) untuk mendapatkan jumlah data di map
	// map[key] untuk mengambil data map dengan key
	// map[key]=value untuk mengubah data dimap dengan key
	// make(map[TypeKey]TypeValue) untuk membuat map baru
	// delete(map,key) unutk menghapus data di map dengan key

	learn := make(map[string]string) // membuat map dengan kata kunci make
	learn["title"] = "Go-Lang"
	learn["author"] = "masendhy"
	learn["ups"] = "salah"

	fmt.Println(learn)
	fmt.Println(len(learn))
	fmt.Println(learn["title"])

	learn["title"] = "Javascript"
	fmt.Println(learn)
	delete(learn, "ups")
	fmt.Println(learn)

}
