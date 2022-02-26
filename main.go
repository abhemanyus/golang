package main

const PI = 3.14

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	width, height float64
}

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.width + rect.height)
}

func (sqr Square) Perimeter() float64 {
	return 4 * sqr.side
}

func (circ Circle) Perimeter() float64 {
	return 2 * PI * circ.radius
}

func (sqr Square) Area() float64 {
	return sqr.side * sqr.side
}

func (rect Rectangle) Area() float64 {
	return rect.height * rect.width
}

func (circ Circle) Area() float64 {
	return PI * circ.radius * circ.radius
}
