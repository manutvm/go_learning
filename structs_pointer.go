package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func (point *Point) Move(dx int, dy int) {
	point.x += dx
	point.y += dy
}

func main() {
	pointer := &Point{
		x: 1,
		y: 2,
	}

	fmt.Println(*pointer)
	pointer.Move(2, 4)

	fmt.Println(*pointer)
}
