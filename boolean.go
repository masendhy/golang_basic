package main

import "fmt"

func main() {
	fmt.Println("Benar = ", true)
	fmt.Println("Salah = ", false)

	// boolean operator

	var ujian = 70
	var absensi = 80

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80

	var lulus = lulusUjian && lulusAbsensi
	fmt.Println(lulus)
}
