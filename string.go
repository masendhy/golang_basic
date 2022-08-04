package main

import "fmt"

// cara menulis string
func main() {
	fmt.Println("masendhy")
	fmt.Println("masendhy", "boedhi", "satriya")

	// function dalam string

	fmt.Println(len("masendhy"))                          // untuk menghitung panjang string (character) = 8
	fmt.Println("masendhy"[1])                            // untuk mengambil  data pada array ke-1 dalam bentuk byte = 97
	fmt.Println("masendhy"[1], "boedhi"[1], "setriya"[1]) // untuk mengambil  data pada array ke-5 dalam bentuk byte = 97 111 101
}
