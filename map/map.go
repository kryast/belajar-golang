package main

import "fmt"

func main() {
	var person = make(map[string]string) // cara pertama membuat map, make(map[tipe data]tipe data)

	person["Nama"] = "Ahmad Syarifuddin"
	person["alamat"] = "Palembang"

	fmt.Println(person)

	person2 := map[string]string{ // cara kedua membuat map, map[tipe data]tipe data { }
		"Nama":   "Ahmad Syarifuddin",
		"alamat": "Palembang",
	}

	fmt.Println(person2)

	person["skill"] = "back-end development" // menambahkan key: value baru kedalam map person

	fmt.Println(person)

	delete(person, "skill") // menghapus key: value di dalam map, delete(map, key)
	fmt.Println(person)
	fmt.Println(len(person)) // menghitung panjang map
}
