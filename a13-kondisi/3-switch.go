package main

import "fmt"

// Switch merupakan seleksi kondisi yang sifatnya fokus
// pada satu variabel, lalu kemudian di-cek nilainya.

func main() {
	var point = 8

	switch point {
	case 8:
		fmt.Println("perfect")
		// Keyword fallthrough digunakan untuk memaksa
		// proses pengecekan diteruskan ke satu case selanjutnya
		// dengan tanpa menghiraukan nilai kondisinya
		fallthrough
	// multiple conditions (separated by comma)
	case 7, 6, 5, 4: {
		fmt.Println("awesome")
	}
	// tanda kurawal ({}) bersifat opsional
	default: {
		fmt.Println("not bad")
		fmt.Println("you can do better!")
	}
	}

	var day = "sunday"

	// switch gaya ala if else
	switch {
	case day == "sunday": {
		fmt.Println("today is sunday")
	}
	case day == "friday": {
		fmt.Println("today is friday")
	}
	}
}
