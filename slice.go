package main

import "fmt"

func main() {
	var days = [7]string{
		"Monday",
		"Teusday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	var slice1 = days[2:5]
	fmt.Println(slice1)

	// Function SLice
	// len(slice) untuk mendapatkan data panjang array slice
	// cap(slice) untuk mendapatkan kapasitas
	// append(slice, data) untuk membuat slicebaru dengan menambah data ke posisi terakhir slice, jika kapasitas udah penuh maka akan membuat arrray baru
	// make([]TypeData, length,capacity) untuk membuat slice baru
	// copy(destination, source) untuk menyalin slice dari source ke destination

	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	var slice2 = days[5:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "January")
	fmt.Println(slice3)
	fmt.Println(days)

	newSlice := make([]string, 3, 5)
	newSlice[1] = "February"
	newSlice[2] = "March"

	fmt.Println(newSlice)

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
