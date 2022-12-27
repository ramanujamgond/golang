// structures in go

package main

import (
	"fmt"
)

func main() {
	// rect1 := Rectange{height: 10, width: 5}
	rect1 := Rectange{10, 5}

	fmt.Println(rect1.height)
	fmt.Println(rect1.width)

	fmt.Println("Area of rectange is", rect1.area())
}

type Rectange struct {
	height float64
	width  float64
}

func (rect *Rectange) area() float64 {
	return rect.height * rect.width
}
