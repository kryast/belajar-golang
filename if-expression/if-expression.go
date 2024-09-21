package main

import "fmt"

func main() {
	name := "Ahmad"

	if name == "Ahmad" {
		fmt.Println("hallo", name)
	} else if name == "Syarifuddin" {
		fmt.Println("hallo", name)
	} else {
		fmt.Println("hallo")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
