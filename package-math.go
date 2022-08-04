package main

import (
	"fmt"
	"math"
)

func main() {
	//math.Round(float64) : membulatkan float64 ke atas atau kebawaah, sesuai yang paling dekat
	x := math.Round(10.5)
	fmt.Println(x)

	//math.FLoor(float64) : Dipaksa pembulatan ke bawah
	y := math.Floor(10.9)
	fmt.Println(y)

	//math.Ceil(float64) : Dipaksa pembulatan ke atas
	s := math.Ceil(10.3)
	fmt.Println(s)

	//math.Max(float64, float64) : Mengembalikan nilai float64 terbesar
	m := math.Max(10, 20)
	fmt.Println(m)

	//math.Min(x, y float64) : Mengembalikan nilai float64 terkecil
	v := math.Min(10, 20)
	fmt.Println(v)

}
