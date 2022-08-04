package main

import "fmt"

func main() {

	counter := 1

	for counter <= 5 {
		fmt.Println("Perulangan ke", counter)
		counter++

	}

	// atau dengan cara

	for counter := 1; counter <= 7; counter++ {
		fmt.Println("Perulangan ke", counter)
	}

	// akses array / slice menggunakan perulangan

	slice := []string{"sendhy", "boedhi", "satriya", "ghaza", "raffa", "axa"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// akses array / slice menggunakan for range

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}
	// i bisa di ganti dengan _ jika tidak di deklarasikan
	for _, value := range slice {
		fmt.Println(value)
	}

	// akses map menggunakan for  range

	person := map[string]string{
		"name":  "sendhy",
		"title": "programmer",
	}

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}
