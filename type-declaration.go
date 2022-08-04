package main

import "fmt"

func main() {
	type noKTP string
	type Married bool

	var noKtpAyah noKTP = "1234567890"
	var marriedStatus Married = true

	fmt.Println(noKtpAyah)
	fmt.Println((marriedStatus))

}
