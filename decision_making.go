package main

import "fmt"

func main() {
	var age int = 30

	if age > 18 {
		fmt.Println("Yes you can vote")
	} else {
		fmt.Println("No you can't vote")
	}
}
