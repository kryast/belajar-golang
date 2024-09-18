package main

import "fmt"

func main() {
	var nilaiUAS = 90
	var absensi = 80

	var lulus bool = nilaiUAS >= 80 && absensi >= 75 // && hasilnya true jika a dan b keduanya true
	var lulus2 bool = nilaiUAS > 80 || absensi > 90  // || hasil nya true jika salah satu dari a atau b hasilnya true

	fmt.Println(lulus)
	fmt.Println(lulus2)
}
