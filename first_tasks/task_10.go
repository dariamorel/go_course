package main

import "fmt"

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	var width, height float64
	fmt.Scan(&width, &height)
	var r = Rectangle{width: width, height: height}
	fmt.Println(r.area())
}
