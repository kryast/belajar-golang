package main

import "fmt"

func main() {
	var nama [2]string // mendeklarasikan array harus ditentukan terlebih dahulu jumlah data di dalam array cth nya [2] yaitu ber isikan 2 jumlah data di dalamnya, lalu di ikuti dengan tipe data di dalam array tersebut

	nama[0] = "Ahmad"
	nama[1] = "Syarifuddin"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[0], nama[1])

	var nilai = [3]int{ // mendeklarasikan array dengan value nya secara langsung
		80,
		90,
		100,
	}

	fmt.Println(nilai[0])
	fmt.Println(nilai[1])
	fmt.Println(nilai[2])

	fmt.Println(len(nama)) // len () digunakan untuk mengukur panjang array
	fmt.Println(len(nilai))

}
