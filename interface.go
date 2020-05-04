package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (circle *Circle) Area() float64 {
	return math.Pi * circle.Radius
}

type Square struct {
	Length  float64
	Breadth float64
}

func (square *Square) Area() float64 {
	return square.Length * square.Breadth
}

type Shapes interface {
	Area() float64
}

func sumAreas(shapes []Shapes) float64 {
	area := 0.0
	for _, shape := range shapes {
		area += shape.Area()
	}
	return area
}

func main() {
	circle := &Circle{12}
	area_circle := circle.Area()

	square := &Square{10, 4}
	area_square := square.Area()

	fmt.Println(area_circle)
	fmt.Println(area_square)

	shapes := []Shapes{circle, square}

	area := sumAreas(shapes)
	fmt.Println(area)
}
