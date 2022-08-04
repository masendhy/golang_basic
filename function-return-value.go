package main

import "fmt"

// membuat function yang dapat mengembalikan data yaitu dengan cara menuliskan type data nya setelah () parameter dan tambahkan kata kunci return untuk mengembalikan datanya

func getHello(name string) string {
	if name == "" {
		return ("Hello bro")
	} else {
		return ("Hello" + " " + name)
	}
}

func main() {
	fmt.Println(getHello("sendhy"))

	result := getHello("boedhi")
	fmt.Println(result)
}
