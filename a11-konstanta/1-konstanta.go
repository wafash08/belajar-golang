package main

import "fmt"

// Konstanta adalah jenis variabel yang nilainya tidak bisa diubah.
// Inisialisasi nilai hanya dilakukan sekali di awal,
// setelah itu variabel tidak bisa diubah nilainya.

func main() {
	const firstName string = "John"
	fmt.Print("Halo ", firstName, "!\n")

	const lastName string = "Wick"
	fmt.Print("Nice to meet you ", firstName ," ", lastName, "!\n")

	const (
		square	string	= "kotak"
		isToday bool	= true
		numeric uint8	= 1
		floatNum	= 2.2
	)

	fmt.Println("========")

	fmt.Println(square)
	fmt.Println(isToday)
	fmt.Println(floatNum)

	const (
		a string	= "konstanta"
		b int		= 3
		c
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("========")

	const (
		today string	= "senin"
		sekarang
		isToday2	= true
	)
	fmt.Println(today)
	fmt.Println(sekarang)
	fmt.Println(isToday)
	fmt.Println("========")

	const satu, dua = 1, 2
	const three, four string = "tiga", "empat"

	fmt.Println(satu, dua)
	fmt.Println(three, four)
}
