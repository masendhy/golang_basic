package main

import "fmt"

// kata kunci sumAll untuk menjumlahkan semua data
func sumAll(numbers ...int) int {

	total := 0
	for _, value := range numbers {
		total += value
	}

	return total

}

func main() {
	total := sumAll(10, 20, 30)
	fmt.Println(total)

	// jika sudah memiliki slice maka data slice dapat dibuat menjadi function variadic dengan merubah slice tersebut menjadi variable arguments yaitu dengan car menambahkan ... pada slice nya

	slice := []int{10, 10, 10}
	total = sumAll(slice...)
	fmt.Println(total)
}
