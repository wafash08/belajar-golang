package main

import "fmt"

func main() {
	var fruits = [4]string{"apple", "banana", "cherry", "delima"}

	// for i := 0; i < len(fruits); i++ {
	// 	fmt.Printf("elemen %d : %s\n", i, fruits[i])
	// }

	// for i, fruit := range fruits {
	// 	fmt.Printf("elemen %d : %s\n", i, fruit)
	// }

	// for _, fruit := range fruits {
	// 	fmt.Printf("elemen : %s\n", fruit)
	// }

	// atau hanya ingin i nya saja
	for i := range fruits {
		fmt.Printf("elemen : %d\n", i)
	}

	var makeFruits = make([]string, 2)
	makeFruits[0] = "mango"
	makeFruits[1] = "orange"

	fmt.Println(makeFruits)
}
