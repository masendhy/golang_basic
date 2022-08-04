package main

import (
	"fmt"
	"os"
)

func main() {

	// menggunakan function .Args untuk mengambil data argumentnya
	args := os.Args
	fmt.Println("Arguments:")
	fmt.Println(args)

	// function .Hostname() unutk mengetahui sistem operasi device kita

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname :", name)
	} else {
		fmt.Println("Error", err.Error())
	}
}
