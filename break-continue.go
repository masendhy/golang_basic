package main

import "fmt"

func main() {

	// break digunakan untuk menghentikan pemrograman

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 5 {
			break
		}
	}

	//kata kunci continue digunakan jika kondisi terpenuhi maka data akan di skip/ lewati

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("Angka Ganjil", "=", i)
	}
}
