package main

import "fmt"

// Variabel temporary adalah variabel yang
// hanya bisa digunakan pada blok seleksi kondisi
//  di mana ia ditempatkan saja.

/*
@see https://dasarpemrogramangolang.novalagung.com/A-seleksi-kondisi.html
*/

func main() {
	var point = 8840.0

	if percent := point / 100; percent >= 100 {
		fmt.Printf("%.1f%s perfect!\n", percent, "%")
	} else if percent >= 70 {
		fmt.Printf("%.1f%s good!\n", percent, "%")
	} else {
		fmt.Printf("%.1f%s not bad!\n", percent, "%")
	}
}
