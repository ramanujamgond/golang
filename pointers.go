// pointers in go
package main

import "fmt"

func main() {
	var x int = 5
	fmt.Println("x: ", x)
	fmt.Println("&x", &x)

	// usage of pointers change the value of y in the memory address
	y := 10
	changeValue(&y)
	fmt.Println(y)
}

func changeValue(y *int) {
	*y = 7
}
