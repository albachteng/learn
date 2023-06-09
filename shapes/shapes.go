package shapes

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base, Height float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

func (t Triangle) Area() float64 {
	return .5 * t.Base * t.Height
}

// assumes the apex is within the bounds of the base
func (t Triangle) Perimeter() float64 {
	return t.Base + math.Sqrt(t.Base*t.Base+4*t.Height*t.Height)
}
