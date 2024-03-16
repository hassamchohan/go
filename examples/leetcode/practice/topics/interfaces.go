package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Polygon interface {
	GetSides() int
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	length float64
	width  float64
	Sides  int
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	return (2 * r.length) + (2 * r.width)
}

func (r Rectangle) GetSides() int {
	return r.Sides
}

func Calculate(s Shape) {
	if rect, ok := s.(Polygon); ok {
		fmt.Println(rect.GetSides())
	}
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}

func main() {
	c := Circle{3.45}
	r := Rectangle{7.87, 9.33, 2}

	Calculate(c)
	Calculate(r)

}
