package main

import (
	"fmt"
	"strings"
)

func main() {
	// strings.ToTitle(string) : menjadikan sring menjadi huruf besar semua
	fmt.Println(strings.ToTitle("sendhy"))

	// strings.Trim(string,cutset) : untuk memotong cutset di awal dan akhir string
	fmt.Println(strings.Trim("    masendhy ku  ", " "))

	//strings.ToLower(string) : Membuat semua karakter string menjadi lower case
	fmt.Println(strings.ToLower("SENDHY BOEDHI"))

	// strings.ToUpper(string) : Membuat semua karakter string menjadi upper case
	fmt.Println(strings.ToUpper("sendhy boedhi"))

	// strings.Split(string, separator) : untuk memotong string berdasrkan separator
	fmt.Println(strings.Split("Sendhy Boedhi Satriya", ""))

	//strings.Contains(string,search) : Untuk mengecek apakah string mengandung string lain
	fmt.Println(strings.Contains("sendhy", "y")) // huruf "y" terdapat dlm string "sendhy", maka hasilnya : true
	fmt.Println(strings.Contains("sendhy", "b")) //	huruf "b" terdapat dlm string "sendhy", maka hasilnya : false

	//strings.ReplaceAll(string, from, to) : Untuk mengubah semua string dari from ke to
	fmt.Println(strings.ReplaceAll(" sendhy sendhy boedhi", "sendhy", "masendhy"))

}
