package main

import "fmt"

// function dapat di simpan dalam sebuah parameter
func getGoobye(name string) string {
	return "Goodbye" + " " + name
}

func main() {
	goodBye := getGoobye
	fmt.Println(goodBye("masendhy"))

}
