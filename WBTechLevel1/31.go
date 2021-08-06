package main

import (
	"fmt"
	"math"
)

// Написать программу нахождения расстояния между 2 точками, которые представление в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x int
	y int
}

func NewPoint(x, y int) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func main() {
	point1 := NewPoint(5, 10)
	point2 := NewPoint(10, 12)
	fmt.Println(DistancePoint(point1, point2))
}

func DistancePoint(p1 *Point, p2 *Point) float64 {
	x := float64(p1.x - p2.x)
	y := float64(p1.y - p2.y)
	return math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
}
