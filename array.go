package main

import "fmt"

func main() {

	var names [3]string

	names[0] = "sendhy"
	names[1] = "boedhi"
	names[2] = "satriya"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//deklarasi array dg cara ke-2

	var values = [3]int{

		90,
		80,
		70,
	}

	fmt.Println(values)

	// Function Array
	// len(array) untuk mendapatkan panjang array
	// array[index] untuk mendapakan data pada posisi index tsb
	// array[index]=value untuk mengubah data pada index tsb

}
