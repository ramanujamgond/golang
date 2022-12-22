package main

import "fmt"

func main() {
	age := 16

	switch age {
	case 18:
		fmt.Println("Prepare for college")
		break
	case 16:
		fmt.Println("Dont run after girls")
		break
	case 20:
		fmt.Println("Ger yourselft a job!")
		break
	default:
		fmt.Println("Are you even alive?")
	}
}
