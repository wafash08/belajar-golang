package main

import "fmt"

func main() {
	var numbers1 = [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	var numbers2 = [2][3]int{{1, 2, 3}, {4, 5, 6}}

	fmt.Println("numbers1 :", numbers1)
	fmt.Println("numbers2 :", numbers2)
}
