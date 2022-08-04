package main

import "fmt"

func main() {

	var a = 5
	var b = 7
	var result = a + b

	fmt.Println(result)

	// 	 Argument Assigment
	// a = a + 10  sama dengan a += 10
	// a = a - 10  sama dengan a -= 10
	// a = a * 10  sama dengan a *= 10
	// a = a / 10  sama dengan a /= 10
	// a = a % 10  sama dengan a %= 10

	var i = 10
	i += 10

	var x = 36
	x /= 6

	fmt.Println(i)
	fmt.Println(x)

	// Unary Operator
	// ++ sama dengan a = a + 1
	// -- sama dengan a = a - 1
	// tanda - sama dengan Negative
	// tanda + sama dengan Positive
	// tanda ! sama dengan Boolean kebalikan( tidak)

}
