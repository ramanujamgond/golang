// handle file input output in go lang

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Create("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("Hi, my name is Ramanujam and this file was created using GO!")
	file.Close()

	stream, err := ioutil.ReadFile("sample.txt")

	if err != nil {
		log.Fatal(err)
	}

	s1 := string(stream)

	fmt.Println(s1)
}
