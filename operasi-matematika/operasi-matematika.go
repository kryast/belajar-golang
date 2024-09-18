package main

import "fmt"

func main() {
	var x = 10 // + tambah, - kurang, * kali, / bagi, % modulus
	var y = 2
	x += 10 // x += 10 itu sama saja seperti x = x + 10
	x *= 2

	y++ // y++ itu sama seperti y = y + 1
	y--
	var z = x / y

	fmt.Println(z)

}
