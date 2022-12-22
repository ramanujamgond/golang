package main

import "fmt"

func main() {
	var EvenNum [5]int
	EvenNum[0] = 0
	EvenNum[1] = 1
	EvenNum[2] = 2
	EvenNum[3] = 3
	EvenNum[4] = 4

	fmt.Println(EvenNum)
	fmt.Println(EvenNum[4])
	fmt.Println(EvenNum[3])
	fmt.Println(EvenNum[2])
	fmt.Println(EvenNum[1])
	fmt.Println(EvenNum[0])

	// another way of initialize the array
	NumValue := [6]int{1, 2, 3, 4, 5, 6}

	fmt.Println(NumValue[0])

	for _, value := range NumValue {
		fmt.Println(value)
	}

	// or
	for i, value := range NumValue {
		fmt.Println(value, i)
	}

	// slice array

	numSlice := []int{5, 4, 3, 3, 1, 8, 7}
	sliced := numSlice[3:5]
	fmt.Println(sliced)

	// print array value in reverse order
	intSlice := []int{5, 4, 3, 3, 1, 8, 7}
	// intSliced := intSlice[0:]
	// or
	intSliced := intSlice[:7]
	fmt.Println(intSliced)

	intSlice2 := make([]int, 7)
	copy(intSlice2, intSlice)
	fmt.Println("intSlice2", intSlice2)
	fmt.Println("intSlice", intSlice)

	// apend value
	intSlice3 := append(intSlice2, 10, 12, 13, -99)
	fmt.Println(intSlice3)
}
