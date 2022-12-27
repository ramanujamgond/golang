// recursion in go
// factorial in go

package main

import "fmt"

func main() {
	num := 6

	fmt.Println(factorial(num))
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}

	return num * factorial(num-1)
}
