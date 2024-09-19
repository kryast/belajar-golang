package main

import "fmt"

func main() {
	// mendeklarasikan array terlebih dahulu. isi jumlah data array, jika belum menentukan jumlah data array maka bisa menggunakan ...
	var month = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = month[4:12]

	fmt.Println(slice1)
	fmt.Println(len(slice1)) // panjang slice
	fmt.Println(cap(slice1)) // kapasitas slice bukan kapasitas array

	var slice2 = append(slice1, "bulan baru") // menambahkan data ke dalam slice
	fmt.Println(slice1)
	fmt.Println(slice2)

	var sliceBaru = make([]string, 2, 5) // make([]tipe data, len, cap)
	sliceBaru[0] = "Ahmad"
	sliceBaru[1] = "Syarifuddin"

	fmt.Println(sliceBaru)
	fmt.Println(len(sliceBaru))
	fmt.Println(cap(sliceBaru))

	copySlice := make([]string, len(sliceBaru), cap(sliceBaru)) // buat slice baru, isinya harus sama dengan slice yg akan di copy

	copy(copySlice, sliceBaru) // copy(slice baru, slice yang akan di copy)

	fmt.Println(copySlice)
	fmt.Println(len(copySlice))
	fmt.Println(cap(copySlice))
}
