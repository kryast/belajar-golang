package main

import "fmt"

func main() {
	type huruf string // type digunakan untuk membuat alias dari tipe data, cth nya tipe data string menjadi huruf
	type benarAtauSalah bool

	var nama huruf = "Ahmad Syarifuddin" // tipe data variable string tetapi menggunakan alias huruf
	var programmer benarAtauSalah = true

	fmt.Println(nama)
	fmt.Println(programmer)
}
