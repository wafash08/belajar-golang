package main

import "fmt"

func main() {
	// for "biasa"
	for i := 0; i < 5; i++ {
		fmt.Println("Angka", i)
	}

	// for hanya dengan kondisi
	var i int = 0

	for i < 5 {
		fmt.Println("Angka (for hanya dengan kondisi)", i)
		i++
	}

	// for tanpa kondisi
	var j int = 0

	for {
		fmt.Println("Angka (for tanpa kondisi)", j)

		j++
		if j == 5 {
			break
		}
	}

	// nested for loop
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}

		fmt.Println()
	}
}
