package main

import (
	"fmt"
	"strconv"
)

func main() {
	//strconv.parseBool(string) : untuk mengubah string ke bool
	boolean, err := strconv.ParseBool("true")
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	//strconv.parseFloat(string) : untuk mengubah string ke float
	float64, err := strconv.ParseFloat("200", 32)
	if err == nil {
		fmt.Println(float64)
	} else {
		fmt.Println(err.Error())
	}

	//strconv.parseInt(string) : untuk mengubah string ke int64
	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	//strconv.FormatBool(bool) : mengubah bool ke string
	v := true
	s := strconv.FormatBool(v)
	fmt.Println("%T, %v\n", s, s)

	//strconv.FormatFloat(float, ...) : mengubah float64 ke string
	//strconv.FormatInt(int,...) : mngubah int64 ke string
	value := strconv.FormatInt(100000, 10)
	fmt.Println(value)
}
