package main

import "testing"

type Test struct {
	shape Shape
	want  float64
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []Test{
		{Rectangle{10, 20}, 60},
		{Square{20}, 80},
		{Circle{50}, 314},
	}
	for _, test := range perimeterTests {
		got := test.shape.Perimeter()
		if got != test.want {
			t.Errorf("want %0.2f, got %0.2f", test.want, got)
		}
	}
}

func TestArea(t *testing.T) {
	perimeterTests := []Test{
		{Rectangle{10, 20}, 200},
		{Square{20}, 400},
		{Circle{10}, 314},
	}
	for _, test := range perimeterTests {
		got := test.shape.Area()
		if got != test.want {
			t.Errorf("want %0.2f, got %0.2f", test.want, got)
		}
	}
}
