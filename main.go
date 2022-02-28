package main

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondsInVector(t time.Time) Point {
	ang := secondsInRadians(t)
	return Point{math.Sin(ang), math.Cos(ang)}
}

func secondHandHead(t time.Time, length float64, tail Point) Point {
	vec := secondsInVector(t)
	return Point{tail.X + vec.X*length, tail.Y - vec.Y*length}
}
