// print line

package main

import "fmt"

func main() {
	var name string = "Ramanujam Gond"
	fmt.Println(len(name))
	fmt.Println(name + " is a chill dude")
	const pi float64 = 3.16412732
	var x int = 5
	isbool := true
	fmt.Printf("%f \n", pi)     // format prints a float number
	fmt.Printf("%.3f\n", pi)    // format print a float number with exactly a three precision
	fmt.Printf("%T \n", name)   // prints type of anything
	fmt.Printf("%t \n", isbool) // this will show a boolean
	fmt.Printf("%d \n", x)      // this will show a integer
	fmt.Printf("%b \n", 35)     // this will show a binary
	fmt.Printf("%c \n", 38)     // this will show a character associated with a key code
	fmt.Printf("%c \n", 38)     // this will show the characters that are related to the ASCII code
	fmt.Printf("%x \n", 12)     // this will show the hex code
	fmt.Printf("%e \n", pi)     // this will show the scientific notation
}
