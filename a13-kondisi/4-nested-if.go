package main

import (
	"fmt"
)

func main() {
	var point = 0

	if point > 7 {
		switch point {
		case 10:
			fmt.Println("perfect!")
		default:
			fmt.Println("nice!")
		}
	} else {
		if point == 5 {
			fmt.Println("not bad")
		} else if point == 3 {
			fmt.Println("keep trying")
		} else {
			if point == 0 {
				fmt.Println("try harder!")
			}
			fmt.Println("you can do it")
		}
	}
}
