package main

import "fmt"

func main() {
	var angka32 int32 = 100
	var angka64 int64 = int64(angka32) // konversi dari int 32 ke int 64
	var angka8 int8 = int8(angka32)    // konversi dari int 32 ke int 8

	fmt.Println(angka32)
	fmt.Println(angka64)
	fmt.Println(angka8)

	var nama = "Ahmad"
	var karakter = nama[0] // output nya 65 berupa tipe byte

	fmt.Println(karakter)

	var hurufKarakter = string(karakter) // konversi dari byte ke string

	fmt.Println(hurufKarakter)
}
