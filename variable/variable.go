package main

import "fmt"

func main() {
	var nama string // cara mendeklarasikan variable, ketik var lalu nama variable lalu tipe data nya

	nama = "Ahmad Syarifuddin"
	fmt.Println(nama)

	nama = "Ahmad Syarifuddin yang kedua" // ini akan menimpa value dari variable nama yang sebelumnya
	fmt.Println(nama)

	var namaLengkap = "Ahmad Syarifuddin" // mendeklarasikan variable dengan langsung memasukkan value, tidak perlu menuliskan tipe data setelah nama variable
	fmt.Println(namaLengkap)

	namaProgrammer := "Ahmad Syarifuddin" // mendeklarasikan variable tanpa menulis var sebelum nama variable dan tipe data setelah nama variable, tetapi menggunakan :=
	fmt.Println(namaProgrammer)

	var ( // mendeklarasikan variable lebih dari satu
		namaDepan    = "Ahmad"
		namaBelakang = "Syarifuddin"
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}
