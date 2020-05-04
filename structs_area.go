package main

import "fmt"

type Point struct {
	X int
	Y int
}

func NewPoint(x int, y int) (*Point, error) {
	return &Point{
		X: x,
		Y: y,
	}, nil
}

func main() {
	point, _ := NewPoint(2, 3)
	fmt.Println(point)
}
