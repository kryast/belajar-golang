package main

import "fmt"

func main() {
	const namaDepan = "Ahmad" // constant adalah variable yang value tetap dan tidak bisa ubah
	const namaBelakang = "Syarifuddin"

	fmt.Println(namaDepan, namaBelakang)

	namaDepan = "Tidak Bisa Diubah"
	namaBelakang = "Tidak Bisa Diubah"

}
