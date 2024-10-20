package main

import (
	"fmt"
	"math"
)

type Point struct{ x, y int }

type Cut struct {
	point1, point2 Point
}

func (c Cut) length() float64 {
	var x1 float64 = float64(c.point1.x)
	var x2 float64 = float64(c.point2.x)
	var y1 float64 = float64(c.point1.y)
	var y2 float64 = float64(c.point2.y)
	var l float64 = math.Abs(x1 - x2)
	var w float64 = math.Abs(y1 - y2)
	return math.Sqrt(math.Pow(l, 2) + math.Pow(w, 2))
}

type Triangle struct{ point1, point2, point3 Point }

func (t Triangle) area() float64 {
	var x1 float64 = float64(t.point1.x)
	var x2 float64 = float64(t.point2.x)
	var x3 float64 = float64(t.point3.x)
	var y1 float64 = float64(t.point1.y)
	var y2 float64 = float64(t.point2.y)
	var y3 float64 = float64(t.point3.y)
	return 0.5 * math.Abs(math.Abs(x1-x3)*math.Abs(y2-y3)-math.Abs(x2-x3)*math.Abs(y1-y3))
}

type Circle struct {
	center Point
	radius int
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}

type Shape interface {
	area() float64
}

func printArea(s Shape) {
	result := s.area()
	fmt.Printf("Площадь фигуры: %.2f\n", result)
}

func main() {
	var c Shape = Circle{center: Point{1, 1}, radius: 5}
	var t Shape = Triangle{Point{2, 2}, Point{2, 5}, Point{5, 2}}
	printArea(c)
	printArea(t)
}
