package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// exit control loop
	j := 2
	for j < 8 {
		fmt.Println(j)
		j++
	}

	// nested loop
	for a := 1; a <= 10; a++ {
		for b := 1; b <= a; b++ {
			fmt.Print("*") // it prints the value as it is inside the parenthesis without adding any new line
			// fmt.Printf("*") // can use Printf also
		}
		fmt.Println()
	}
}
