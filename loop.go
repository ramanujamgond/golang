package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// // exit control loop
	// j := 2
	// for j < 8 {
	// 	fmt.Println(j)
	// 	j++
	// }

	// nested loop
	for a := 1; a <= 2; a++ {
		for b := 1; b <= a; b++ {
			fmt.Println("*")
		}
	}
}
