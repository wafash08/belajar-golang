package main

import "fmt"

func main() {
	// cara inisialisasi pertama
	var names [4]string
	names[0] = "a"
	names[1] = "b"
	names[2] = "c"
	names[3] = "d"
	// inisialisasi di bawah ini tidak akan bisa
	// names[4] = "e"

	fmt.Println(names[0], names[1], names[2], names[3])

	// cara inisialisasi kedua

	// secara horizontal
	// var fruits = [4]string{"apple", "banana", "cherry", "delima"}

	// secara vertikal
	var fruits = [4]string{
		"apple",
		"banana",
		"cherry",
		"delima",
	}

	fmt.Println("Jumlah elemen \t\t:", len(fruits))
	// biasa dipakai untuk kegiatan debugging
	fmt.Println("Isi semua elemen \t:", fruits)

	// cara inisialisasi ketiga
	// tanpa jumlah elemen
	var numbers = [...]int{2,3,4,5,6}
	fmt.Println("data array \t\t:", numbers)
	fmt.Println("jumlah elemen \t\t:", len(numbers))
}
