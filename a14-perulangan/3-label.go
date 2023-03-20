package main

import "fmt"

func main() {
	outerloop:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 3 {
				// lanjut ke perulangan selanjutnya, i = 4
				// continue outerloop

				// hentikan perulangan, cukup sampai sini saja, i = 2 saja
				break outerloop
			}

			fmt.Print("matriks [", i, "][", j, "]", "\n")
		}
	}
}
