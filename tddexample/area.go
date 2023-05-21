package tddexample

import "math"

type Shape interface {
	Perimeter() float64
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Circle struct {
	Radius float64
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

func (c Circle) Area() float64 {
	return math.Pi * math.Abs(c.Radius)
}
