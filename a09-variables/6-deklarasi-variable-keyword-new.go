package main

import "fmt"

func main() {
	name := new(string)

	// Variabel name menampung data bertipe pointer string. Jika ditampilkan yang muncul bukanlah nilainya melainkan alamat memori nilai tersebut (dalam bentuk notasi heksadesimal).
	// Untuk menampilkan nilai aslinya, variabel tersebut perlu di-dereference terlebih dahulu, menggunakan tanda asterisk (*).
	fmt.Println(name)  // 0x20818a220
	fmt.Println(*name) // ""
}
