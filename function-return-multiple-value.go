package main

import "fmt"

// membuat multiple value dengan cara menambahkan type data di dalam tanda ()
func getFullName() (string, string, string) {
	return "sendhy", "boedhi", "satriya"
}

func main() {

	firstName, _, lastName := getFullName()

	// untuk meng ignore variable nya dapat digunakan tanda _ (underscore)

	fmt.Println(firstName)
	//fmt.Println(middleName)
	fmt.Println(lastName)
}
